var HOST="/api",Ok=1,title=void 0,type=void 0,created=void 0,update=void 0,cover=void 0,desc=void 0,tags=void 0,keyword=void 0,isTop=void 0,hot=void 0,recommend=void 0,coverStr=void 0,content=window.content;function articleItemValue(){title=$("#article-title").val(),$("#article-type").val()&&""!==$("#article-type").val()&&(type=$("#article-type").val().split(",").map(function(e){return Number(e)})),created=$("#article-created").val(),update=$("#article-update").val(),(desc=$("#article-desc").val())&&"string"==typeof desc&&""!==desc&&54<desc.length&&(desc=desc.substr(0,54)),cover=$(".cupload-image-box input").val(),keyword=$("#keyword").val(),isTop=$("#is-top").prop("checked"),hot=$("#hot").prop("checked"),recommend=$("#recommend").prop("checked")}function save(){var e=0<arguments.length&&void 0!==arguments[0]&&arguments[0];articleItemValue();for(var t=$("#articleTags li"),i=[],o=0;o<t.length;o++)$(t[o]).find(".tagit-label")&&0<$(t[o]).find("span.tagit-label").length&&$(t[o]).find(".tagit-label")[0].innerText&&""!==$(t[o]).find(".tagit-label")[0].innerText&&"string"==typeof $(t[o]).find(".tagit-label")[0].innerText&&i.push($(t[o]).find(".tagit-label")[0].innerText);if(0===i.length?i="":i.join(","),cover)if(get()){var n=get().secretKey;coverStr=getTokenUrl($('.cupload-image-list li input[type="hidden"]').val(),n)}else{var r=handleToken();if(r&&2===Object.keys(r).length){var a=c.secretKey;coverStr=getTokenUrl($('.cupload-image-list li input[type="hidden"]').val(),a)}}var c={title:title,created:created,update:update,content:window.content,cover:coverStr,desc:desc,keyword:keyword||"",isTop:isTop||!1,hot:hot||!1,recommend:recommend||!1,prod:e};""!==i&&(c.Tags=i),void 0!==desc&&""!==desc&&desc||(c.desc=window.text?window.text.substr(0,54):null),type&&Array.isArray(type)&&1===type.length&&(c.menuId=type[0],c.categoryId=-1),type&&Array.isArray(type)&&2===type.length&&(c.menuId=type[0],c.categoryId=type[1]),Object.keys(c).map(function(e){""!==c[e]&&void 0!==c[e]&&null!==c[e]||"keyword"!==e&&delete c[e]}),saveArticle(c)}function saveArticle(e){$.ajax({url:HOST+"/article/details/add",data:JSON.stringify(e),method:"POST",success:function(e){e.code===Ok?(remove(),window.message.success(e),setTimeout(function(){window.location.reload()},5e3)):window.message.error(e)}})}function remove(){cover&&(window.qiniuyun=null),coverStr=recommend=hot=isTop=keyword=tags=desc=cover=update=created=type=title=null,window.content=null,window.text=null}function getTokenUrl(e,t){var i=void 0,o=window.qiniuyun.size,n=window.btoa("0612_"+(new Date).getTime()+"_"+parseInt(10*Math.random())),r=e.split("base64,")[1],a="http://up-z2.qiniup.com/putb64/"+o+"/key/"+n;return $.ajax({url:a,type:"POST",async:!1,beforeSend:function(e){e.setRequestHeader("Content-Type","application/octet-stream"),e.setRequestHeader("Authorization","UpToken "+t)},data:r,success:function(e){i=e.key}}),i}function handleToken(){var t=void 0;return $.ajax({url:HOST+"/qn/token",method:"get",async:!1,success:function(e){e.code===Ok?set({expireTime:e.expireTime,token:e.token},function(e){t=e}):window.message.error(e)}}),t}function set(e,t,i){if(e&&2===Object.keys(e).length){var o={};o.expireTime=e.expireTime,o.secretKey=e.token,localStorage.setItem("localStorage",JSON.stringify(o)),t&&t(o)}else i&&i()}function get(){var e=localStorage.getItem("localStorage"),t=void 0;if(!e||!isJSON(e))return null;if(e=JSON.parse(e),2===Object.keys(e).length){if(!((new Date).getTime()<e.expireTime+36e5))return localStorage.removeItem("localStorage"),null;t=e}return t}function isJSON(e){if("string"==typeof e)try{return JSON.parse(e),!0}catch(e){return!1}}