var modules = ['service',
  'translate',
  'jQuery',
  'bootstrap'
];

define(modules, function (service,translate, $, bootstrap) {
  $('#botao').on('click', service.postLead);
  bindLangEvents();
  translate.translate(window.location.hash);

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

