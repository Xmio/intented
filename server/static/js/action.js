var modules = ['service',
  'jQuery',
  'bootstrap'
];

define(modules, function (service, $, bootstrap) {

  $('#botao').on('click', service.postLead);

});

