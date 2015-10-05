var modules = ['jQuery','labels'];

define(modules, function ($, labels) {

 function translate(){
    var exports = {};
    exports.translate = translate;

    function translate(hash){
      var lang = hash.toUpperCase();
     if(lang.indexOf("PT") > -1)
        build("PT");
      else if(lang.indexOf("ES") > -1)
        build("ES");
      else
        build("EN");
    }

    function build(language){
      $('.page-tittle').text(labels.get('TITLE', language));
    }
    return exports;
  }

  return translate();

});
