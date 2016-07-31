$( document ).ready(function doPoll() {
  console.log('ready');
  $.getJSON("/attributes/list", function( data ) {
    console.log('making a call');
  var items = [];
  $.each( data, function( key, val ) {
    if (val.Known) {
      var color;
      if (val.Presence) {
        color = "#BBFFBB";
      } else {
        color = "#FFBBBB";
      }
      items.push( "<button class='btn btn-default' id='" + key + "' style='background-color:" + color + "'>" + JSON.stringify(val.Name) + "</button>" );
    }

  });
  $('#attributes').html(items.join( "" ));
  setTimeout(doPoll,5000);
});
});
