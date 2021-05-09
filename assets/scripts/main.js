var editor = document.getElementById("editor");

var pluginBar = document.getElementById("plugin-navber");

pluginBar.innerHTML = loadNavbar();

setInterval(() => {
  editor = document.getElementById("editor");
}, 100);
