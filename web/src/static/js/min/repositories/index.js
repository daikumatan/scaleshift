var images=[],conditions={firstLoad:!0,words:""},vue=new Vue({el:"#data",data:{images:[]},methods:{update:function(){var o=[];$.map(images,function(e){(""==conditions.words||app.match([e.namespace,e.name],conditions.words))&&o.push({code:e.namespace+"/"+e.name,namespace:e.namespace,name:e.name,description:e.description?marked(e.description):"---"})}),o.sort(function(e,o){return e.name<o.name?-1:e.name>o.name?1:0}),this.images=o,conditions.firstLoad&&(conditions.firstLoad=!1,""!=conditions.words&&0<o.length&&setTimeout(function(){$("#query-words").blur()},500)),$("#record-count").text(o.length),$(".wait-icon").hide()}}});function update(){conditions.words=app.singleSpace($("#query-words").val()),vue.update()}function load(a){API("Repository").getRemoteRepositories(function(e,o,t){if(app.shouldExit(t,e))return alert("Something went wrong. Check your configurations!"),void(window.location.href="/settings/");images=t.body,update(),a&&a()})}function loadDetails(n){if(!$(n).attr("data-loaded")){var i=$(n).attr("data-ns");i&&(i+="/");var r=$(n).attr("data-nm");API("Repository").getRemoteImages(r,function(e,o,t){if(app.shouldExit(t,e))return alert("Something went wrong. Check your configurations!"),void(window.location.href="/settings/");t.body.sort(function(e,o){return e.repoTags[0]<o.repoTags[0]?1:e.repoTags[0]>o.repoTags[0]?-1:0});var a="";$.map(t.body,function(e){var o=e.repoTags[0];a+='<tr data-id="'+i+r+":"+o+'">',a+="<td>"+o+"</td>",a+='<td><a class="waves-effect waves-light btn blue darken-1">download</a></td>',a+="</tr>"}),$(n).find(".progress").hide(),$(".row-body tbody",n).html(a),$(".row-body .btn",n).click(function(e){pullImage($(e.target).closest("tr").attr("data-id"))}),$(n).attr("data-loaded","done")})}}function pullImage(a){var e=new models.ImageName(a);API("Image").postNewImage(e,function(e,o,t){if(app.shouldExit(t,e))return alert("Something went wrong. Check your configurations!"),void(window.location.href="/settings/");location.href="/images/?q="+encodeURIComponent(a)})}$(document).ready(function(){config.get().usePrivateRegistry?($("#menu-repositories").addClass("active"),app.query("q")&&$("#query-words").val(app.query("q")).focus(),$(".collapsible").on("shown.bs.collapse",function(e){loadDetails($(e.target).closest("li"))}),load(function(){$("#query-words").keyup(function(){update()}),$("#data").fadeIn()})):window.location.href="/images/"});