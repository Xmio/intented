var modules = ['jQuery'];

define(modules, function ($) {

    function service(){
      var exports = {};
      exports.postLead = postLead;

      function postLead(){
        $.ajax({
          url : '/lead/',
          type : 'POST',
          contentType : "application/json; charset=utf-8",
          success: onSuccess,
          error : showError,
          data : JSON.stringify(getLead())
        });
      }

      function getLead(){

      }

      function showError(){

      }

      function onSuccess(){

      }

      return exports;
    }

    return service();
});

