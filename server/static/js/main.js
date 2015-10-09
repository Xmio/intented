requirejs.config({
  paths: {
        'jQuery' : 'https://ajax.googleapis.com/ajax/libs/jquery/2.1.4/jquery.min',
        'bootstrap' : 'https://maxcdn.bootstrapcdn.com/bootstrap/3.3.5/js/bootstrap.min',
        'notify' : 'https://cdnjs.cloudflare.com/ajax/libs/notify/0.3.2/notify'
    },
    shim: {
        'jQuery': {
            exports: '$'
        },
         'bootstrap': {
            deps: ['jQuery']
        },
         'notify': {
            deps: ['jQuery','bootstrap']
        }
    }
});

require(['action']);
