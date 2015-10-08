var modules = ['jQuery'];

define(modules, function ($) {

    function service(){
      var lead = {
        mail :'',
        invited : ''
      }
      var exports = {};
      exports.postLead = postLead;

      function postLead(){
        $.ajax({
          url : '/lead',
          type : 'POST',
          success: onSuccess,
          data : getLead()
        });
      }

      function onSuccess() {

      }

      function getLead() {
        lead.mail=$(".email")[0].value;
        return lead;
      }

      $(".call-to-lead").click(validate)

      function validate() {
        email =$(".email")[0].value;
        if (!email || email === '') {
          alert("informe o email");
          return false;
        } else
          postLead();
      }

      return exports;
    }

    return service();
});
