dlFontsBtn = document.createElement('button');
dlFontsBtn.id = "font-btn";
dlFontsBtn.innerHTML = "Download Required Fonts";

body = document.getElementsByTagName("body")[0];
body.innerHTML = dlFontsBtn.outerHTML + body.innerHTML;

dlFontsBtn = document.getElementById("font-btn");
dlFontsBtn.onclick = downloadFonts;


/*window.onload = downloadFontsAsEmbedded;
document.fonts.ready.then(embedFonts);

function downloadFontsAsEmbedded() {
    var FiraSansLight = new FontFace('Fira Sans Light Embedded', 'url("fonts/FiraSans-Light.ttf")');
    document.fonts.add(FiraSansLight);
    var FiraCode = new FontFace('Fira Code Embedded', 'url("fonts/FiraCode-Regular.ttf")');
    document.fonts.add(FiraCode);
    var FiraSansSemiBold = new FontFace('Fira Sans SemiBold Embedded', 'url("fonts/FiraSans-SemiBold.ttf")');
    document.fonts.add(FiraSansSemiBold);
    var Merriweather = new FontFace('Merriweather Embedded', 'url("fonts/Merriweather-Regular.ttf")');
    document.fonts.add(Merriweather);
}

function embedFonts() {
    style = document.createElement('style');
    style.type = 'text/css';
    style.innerHTML = `
        html                 { font-family: "Fira Sans Light", "Fira Sans Light Embedded", sans-serif; }
        h1, h2               { font-family: "Fira Sans SemiBold", "Fira Sans SemiBold Embedded", sans-serif; }
        h3, h4, h5, h6       { font-family: "Merriweather", "Merriweather Embedded", serif; }
        div.post-header time { font-family: "Fira Code", "Fira Code Embedded", monospace; }`;
    document.getElementsByTagName('head')[0].appendChild(style);
}*/

function downloadFonts() {


    var FiraSansLight = new FontFace('Fira Sans Light', 'url("fonts/FiraSans-Light.ttf")');
    var FiraCode = new FontFace('Fira Code', 'url("fonts/FiraCode-Regular.ttf")');
    var FiraSansSemiBold = new FontFace('Fira Sans SemiBold', 'url("fonts/FiraSans-SemiBold.ttf")');
    var Merriweather = new FontFace('Merriweather', 'url("fonts/Merriweather-Regular.ttf")');
    document.fonts.add(FiraSansLight);
    document.fonts.add(FiraCode);
    document.fonts.add(FiraSansSemiBold);
    document.fonts.add(Merriweather);


}