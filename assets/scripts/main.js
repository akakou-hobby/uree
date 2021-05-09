var editor = document.getElementById("editor");

var pluginBar = document.getElementById("plugin-navber");
var sideIcons = document.getElementById("side-icons");

pluginBar.innerHTML = loadNavbar();
sideIcons.innerHTML = loadSidePallet();

setInterval(() => {
  editor = document.getElementById("editor");
}, 100);
