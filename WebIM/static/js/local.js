function capitalizeFirstLetter(string) {
  return string.charAt(0).toUpperCase() + string.slice(1);
}
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
      items.push( "<div class='col-xs-6'><button class='btn btn-default' id='" + key + "' style='width:200px;background-color:" + color + "'>" + capitalizeFirstLetter(val.Name) + "</button></div>" );
    }

  });
  $('#attributes').html(items.join( "" ));
  setTimeout(doPoll,5000);
});
});
