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
          success: onSuccessCreateLead,
          data : getLead()
        });
      }

      function onSuccessCreateLead(referal) {
        referalUrl='http://www.newintend.com/?ref='+referal;
        encondedReferalUrl = encodeURIComponent(referalUrl);
        $('.btn-facebook').attr('href','https://www.facebook.com/sharer/sharer.php?u='+encondedReferalUrl);
        $('.btn-google').attr('href','https://plus.google.com/share?url='+encondedReferalUrl);
        $('.btn-linkedin').attr('href','http://www.linkedin.com/shareArticle?title=NewIntend&mini=true&url='+encondedReferalUrl);
        $('.btn-twitter').attr('href','http://twitter.com/home?status='+encondedReferalUrl+' Crie propostas eficientes melhorando a comunicação com seu cliente');
        $('.referal-link').text(referalUrl);
      }

      function getLead() {
        lead.mail=$('.email')[0].value;
        lead.invited=getUrlParameter('ref');
        return lead;
      }

      $('.call-to-lead').click(validate)

      function validate() {
        email =$('.email')[0].value;
        if (!email || email === '') {
          alert('informe o email');
          return false;
        } else
          postLead();
      }

      function getUrlParameter(param) {
        var sPageURL = decodeURIComponent(window.location.search.substring(1));
        var sURLVariables = sPageURL.split('&');
        for (i = 0; i < sURLVariables.length; i++) {
            sParameterName = sURLVariables[i].split('=');
            if (sParameterName[0] === param)
                return sParameterName[1] === undefined ? true : sParameterName[1];
        }
        return '';
      };

      $(document).ready(function() {
        $('.btn-social-icon').click(function(e) {
            e.preventDefault();
            window.open($(this).attr('href'), 'fbShareWindow', 'height=450, width=550, top=' + ($(window).height() / 2 - 275) + ', left=' + ($(window).width() / 2 - 225) + ', toolbar=0, location=0, menubar=0, directories=0, scrollbars=0');
            return false;
        });
    });

      return exports;
    }

    return service();
});
