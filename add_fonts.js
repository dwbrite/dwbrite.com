var clicked = false;

if(sessionStorage.getItem("cachedFonts") ==="true") {
    downloadFonts();
}

var dlFontsBtn = document.createElement('button');
dlFontsBtn.id = "font-btn";
dlFontsBtn.title = `There's no good way to make sure that you have the fonts, 
so I guess I'll have to keep this button here forever.`;
dlFontsBtn.innerHTML = "Click for Fonts";

var body = document.getElementsByTagName("body")[0];
body.innerHTML = dlFontsBtn.outerHTML + body.innerHTML;

dlFontsBtn = document.getElementById("font-btn");
dlFontsBtn.onclick = downloadFonts;

function downloadFonts() {
    if(!clicked) {
        addFontLink('fonts/Fira/fira.css');
        addFontLink('fonts/FiraCode/fira_code.css');
        addFontLink('fonts/Merriweather/merriweather.css');
    }
    clicked = true;
    sessionStorage.setItem("cachedFonts", "true")
}

function addFontLink(location) {
    var head = document.getElementsByTagName('head')[0];
    var linkStyle = document.createElement('link');
    linkStyle.setAttribute('rel', 'stylesheet');
    linkStyle.setAttribute('type', 'text/css');
    linkStyle.setAttribute('href', location);
    head.appendChild(linkStyle);
}