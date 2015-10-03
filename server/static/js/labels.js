var modules = [];

define(modules, function () {

    function labels(){
      var labels = {};
      labels.get = get;

      function get(name, language) {
        return labels[name][language];
      }

      labels.TITLE = {
        EN : 'Big projects start with greate ',
        PT : 'Grandes projetos começam com ótimas negociações',
        ES : 'Grandes ermans'
      };

      return labels;
    }

    return labels();
});
