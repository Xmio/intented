var modules = ['service',
  'translate',
  'jQuery',
  'bootstrap'
];

define(modules, function (service,translate, $, bootstrap) {
  $('#botao').on('click', service.postLead);
  bindLangEvents();
  if(window.location.hash)
    translate.translate(window.location.hash)
  else
    translate.translate(navigator.language);

  function bindLangEvents(){
    $('.pt-lang').on('click', function(){
      translate.translate('PT');
    });
    $('.en-lang').on('click', function(){
      translate.translate('EN');
    });
    $('.es-lang').on('click', function(){
      translate.translate('ES');
    });
  }
});

