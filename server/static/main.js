requirejs.config({
  paths: {
        'jQuery' : 'https://ajax.googleapis.com/ajax/libs/jquery/2.1.4/jquery.min',
        'bootstrap' : 'https://maxcdn.bootstrapcdn.com/bootstrap/3.3.5/js/bootstrap.min'
    },
    shim: {
        'jQuery': {
            exports: '$'
        }
    }
});

require(['service']);
