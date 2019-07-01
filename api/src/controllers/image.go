package controllers

import (
	"context"
	"net/http"
	"strings"
	"time"

	"docker.io/go-docker/api/types"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	"github.com/rescale-labs/scaleshift/api/src/auth"
	"github.com/rescale-labs/scaleshift/api/src/config"
	"github.com/rescale-labs/scaleshift/api/src/db"
	"github.com/rescale-labs/scaleshift/api/src/generated/models"
	"github.com/rescale-labs/scaleshift/api/src/generated/restapi/operations"
	"github.com/rescale-labs/scaleshift/api/src/generated/restapi/operations/image"
	"github.com/rescale-labs/scaleshift/api/src/log"
	"github.com/rescale-labs/scaleshift/api/src/queue"
)

func imageRoute(api *operations.ScaleShiftAPI) {
	api.ImageGetImagesHandler = image.GetImagesHandlerFunc(getImages)
	api.ImagePostNewImageHandler = image.PostNewImageHandlerFunc(postNewImage)
	api.ImageDeleteImageHandler = image.DeleteImageHandlerFunc(deleteImage)
}

func getImages(params image.GetImagesParams) middleware.Responder {
	cli, _, code := dockerClient(nil)
	if code != 0 {
		return image.NewGetImagesDefault(code).WithPayload(newerror(code))
	}
	defer cli.Close()

	// local docker images
	images, err := cli.ImageList(context.Background(), types.ImageListOptions{})
	if err != nil {
		log.Error("ImageList@getImages", err, nil)
		code := http.StatusBadRequest
		return image.NewGetImagesDefault(code).WithPayload(newerror(code))
	}
	result := []*models.Image{}
	for _, image := range images {
		tags := []string{}
		for _, tag := range image.RepoTags {
			appendable := true
			for _, ignore := range config.Config.ImagesToBeIgnored {
				if strings.Contains(strings.ToLower(tag), strings.ToLower(ignore)) {
					appendable = false
				}
			}
			if !appendable {
				continue
			}
			tags = append(tags, tag)
		}
		if len(tags) == 0 {
			continue
		}
		imageName := strings.TrimPrefix(image.ID, "sha256:")
		result = append(result, &models.Image{
			ID:          swag.String(imageName[0:min(12, len(imageName))]),
			ParentID:    image.ParentID,
			RepoDigests: image.RepoDigests,
			RepoTags:    tags,
			Status:      db.StableImage,
			Size:        image.Size,
			VirtualSize: image.VirtualSize,
			Created:     time.Unix(image.Created, 0).Format(time.RFC3339),
		})
	}
	// pulling Images
	if pImages, err := db.GetPullingImages(); err == nil {
		for _, image := range pImages {
			if inStables(image.Tag, images) {
				continue
			}
			result = append(result, &models.Image{
				RepoTags: []string{image.Tag},
				Status:   image.Status,
				Created:  image.Started.Format(time.RFC3339),
			})
		}
	}
	return image.NewGetImagesOK().WithPayload(result)
}

func inStables(candidate string, images []types.ImageSummary) bool {
	for _, image := range images {
		for _, tag := range image.RepoTags {
			if strings.EqualFold(tag, candidate) {
				return true
			}
		}
	}
	return false
}

func postNewImage(params image.PostNewImageParams) middleware.Responder {
	var cfg *types.AuthConfig
	name := swag.StringValue(params.Body.Image)
	if strings.HasPrefix(name, config.Config.DockerRegistryHostName) {
		if sess, err := auth.RetrieveSession(params.HTTPRequest); err == nil && sess != nil {
			creds := auth.FindCredentials(sess.DockerUsername)
			cfg = &types.AuthConfig{
				ServerAddress: config.Config.DockerRegistryHostName,
				Username:      creds.Base.DockerUsername,
				Password:      creds.Base.DockerPassword,
			}
		} else {
			code := http.StatusForbidden
			return image.NewPostNewImageDefault(code).WithPayload(newerror(code))
		}
	}
	if strings.HasPrefix(name, config.Config.NgcRegistryHostName) {
		if sess, err := auth.RetrieveSession(params.HTTPRequest); err == nil && sess != nil {
			creds := auth.FindCredentials(sess.DockerUsername)
			cfg = &types.AuthConfig{
				ServerAddress: config.Config.NgcRegistryHostName,
				Username:      config.Config.NgcRegistryUserName,
				Password:      creds.Base.NgcApikey,
			}
		} else {
			code := http.StatusForbidden
			return image.NewPostNewImageDefault(code).WithPayload(newerror(code))
		}
	}
	cli, auth, code := dockerClient(cfg)
	if code != 0 {
		return image.NewPostNewImageDefault(code).WithPayload(newerror(code))
	}
	defer cli.Close()

	if err := db.SetPullingImageMeta(name); err != nil {
		log.Error("SetPullingImageMeta@postNewImage", err, nil)
		code := http.StatusInternalServerError
		return image.NewPostNewImageDefault(code).WithPayload(newerror(code))
	}
	if err := queue.SubmitPullImageJob(name, swag.StringValue(auth)); err != nil {
		log.Error("SubmitPullImageJob@postNewImage", err, nil)
		code := http.StatusInternalServerError
		return image.NewPostNewImageDefault(code).WithPayload(newerror(code))
	}
	return image.NewPostNewImageCreated()
}

func deleteImage(params image.DeleteImageParams) middleware.Responder {
	cli, _, code := dockerClient(nil)
	if code != 0 {
		return image.NewDeleteImageDefault(code).WithPayload(newerror(code))
	}
	defer cli.Close()

	name := swag.StringValue(params.Body.Image)
	options := types.ImageRemoveOptions{
		Force:         false,
		PruneChildren: true,
	}
	if _, err := cli.ImageRemove(context.Background(), name, options); err != nil {
		log.Error("ImageRemove@deleteImage", err, nil)
		code := http.StatusInternalServerError
		return image.NewDeleteImageDefault(code).WithPayload(newerror(code))
	}
	return image.NewDeleteImageNoContent()
}
