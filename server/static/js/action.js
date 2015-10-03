var modules = ['service',
  'bootstrap',
  'jQuery'
];

define(modules, function (service,bootstrap, $) {

  $('#botao').on('click', service.postLead);

});

