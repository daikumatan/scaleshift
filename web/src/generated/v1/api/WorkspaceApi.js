/**
 * ScaleShift
 * A platform for machine learning & high performance computing 
 *
 * OpenAPI spec version: 1.0.0
 *
 * NOTE: This class is auto generated by the swagger code generator program.
 * https://github.com/swagger-api/swagger-codegen.git
 *
 * Swagger Codegen version: 2.3.1
 *
 * Do not edit the class manually.
 *
 */

(function(root, factory) {
  if (typeof define === 'function' && define.amd) {
    // AMD. Register as an anonymous module.
    define(['ApiClient', 'model/Error', 'model/Workspace'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    module.exports = factory(require('../ApiClient'), require('../model/Error'), require('../model/Workspace'));
  } else {
    // Browser globals (root is window)
    if (!root.ScaleShift) {
      root.ScaleShift = {};
    }
    root.ScaleShift.WorkspaceApi = factory(root.ScaleShift.ApiClient, root.ScaleShift.Error, root.ScaleShift.Workspace);
  }
}(this, function(ApiClient, Error, Workspace) {
  'use strict';

  /**
   * Workspace service.
   * @module api/WorkspaceApi
   * @version 1.0.0
   */

  /**
   * Constructs a new WorkspaceApi. 
   * @alias module:api/WorkspaceApi
   * @class
   * @param {module:ApiClient} [apiClient] Optional API client implementation to use,
   * default to {@link module:ApiClient#instance} if unspecified.
   */
  var exports = function(apiClient) {
    this.apiClient = apiClient || ApiClient.instance;


    /**
     * Callback function to receive the result of the deleteWorkspace operation.
     * @callback module:api/WorkspaceApi~deleteWorkspaceCallback
     * @param {String} error Error message, if any.
     * @param data This operation does not return a value.
     * @param {String} response The complete HTTP response.
     */

    /**
     * delete user&#39;s workspace 
     * @param {module:model/Workspace} body 
     * @param {module:api/WorkspaceApi~deleteWorkspaceCallback} callback The callback function, accepting three arguments: error, data, response
     */
    this.deleteWorkspace = function(body, callback) {
      var postBody = body;

      // verify the required parameter 'body' is set
      if (body === undefined || body === null) {
        throw new Error("Missing the required parameter 'body' when calling deleteWorkspace");
      }


      var pathParams = {
      };
      var queryParams = {
      };
      var collectionQueryParams = {
      };
      var headerParams = {
      };
      var formParams = {
      };

      var authNames = [];
      var contentTypes = ['application/json'];
      var accepts = ['application/json'];
      var returnType = null;

      return this.apiClient.callApi(
        '/workspaces', 'DELETE',
        pathParams, queryParams, collectionQueryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, callback
      );
    }

    /**
     * Callback function to receive the result of the getWorkspaces operation.
     * @callback module:api/WorkspaceApi~getWorkspacesCallback
     * @param {String} error Error message, if any.
     * @param {Array.<module:model/Workspace>} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * returns user&#39;s workspaces 
     * @param {module:api/WorkspaceApi~getWorkspacesCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link Array.<module:model/Workspace>}
     */
    this.getWorkspaces = function(callback) {
      var postBody = null;


      var pathParams = {
      };
      var queryParams = {
      };
      var collectionQueryParams = {
      };
      var headerParams = {
      };
      var formParams = {
      };

      var authNames = [];
      var contentTypes = ['application/json'];
      var accepts = ['application/json'];
      var returnType = [Workspace];

      return this.apiClient.callApi(
        '/workspaces', 'GET',
        pathParams, queryParams, collectionQueryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, callback
      );
    }
  };

  return exports;
}));