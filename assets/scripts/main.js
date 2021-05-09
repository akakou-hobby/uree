var editor = document.getElementById("editor");

var pluginBar = document.getElementById("plugin-navber");
var sideIcons = document.getElementById("side-icons");
var sidePallet = document.getElementById("side-pallet");

pluginBar.innerHTML = loadNavbar();
sideIcons.innerHTML = loadSidePallet();

setInterval(() => {
  editor = document.getElementById("editor");
  sidePallet = document.getElementById("side-pallet");
}, 100);
