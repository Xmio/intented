var modules = ['jQuery','labels'];

define(modules, function ($, labels) {

 function translate(){
    var exports = {};
    exports.translate = translate;

    function translate(hash){
      var lang = hash.substring(1).toUpperCase();
     if(lang == "PT")
        build(lang);
      else if(hash == "ES")
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
