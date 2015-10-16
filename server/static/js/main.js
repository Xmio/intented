requirejs.config({
  paths: {
        'jQuery' : 'https://ajax.googleapis.com/ajax/libs/jquery/2.1.4/jquery.min',
        'bootstrap' : 'https://maxcdn.bootstrapcdn.com/bootstrap/3.3.5/js/bootstrap.min',
        'notification' :'https://cdnjs.cloudflare.com/ajax/libs/jquery.notification/1.0.2/jquery.notification.min'
    },
    shim: {
        'jQuery': {
            exports: '$'
        },
         'bootstrap': {
            deps: ['jQuery']
        },
         'notification': {
            deps: ['jQuery']
        }
    }
});

require(['action']);
