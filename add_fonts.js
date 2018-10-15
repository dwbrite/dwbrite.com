window.onload = function() {
    var FiraSansLight = new FontFace('Fira Sans Light Embedded', 'url("fonts/FiraSans-Light.ttf")');
    document.fonts.add(FiraSansLight);
    var FiraCode = new FontFace('Fira Code Embedded', 'url("fonts/FiraCode-Regular.ttf")');
    document.fonts.add(FiraCode);
    var FiraSansSemiBold = new FontFace('Fira Sans SemiBold Embedded', 'url("fonts/FiraSans-SemiBold.ttf")');
    document.fonts.add(FiraSansSemiBold);
    var Merriweather = new FontFace('Merriweather Embedded', 'url("fonts/Merriweather-Regular.ttf")');
    document.fonts.add(Merriweather);
};

document.fonts.ready.then( function() {
    style = document.createElement('style');
    style.type = 'text/css';
    style.innerHTML = `
        html                 { font-family: "Fira Sans Light", "Fira Sans Light Embedded", sans-serif; }
        h1, h2               { font-family: "Fira Sans SemiBold", "Fira Sans SemiBold Embedded", sans-serif; }
        h3, h4, h5, h6       { font-family: "Merriweather", "Merriweather Embedded", serif; }
        div.post-header time { font-family: "Fira Code", "Fira Code Embedded", monospace; }`;
    document.getElementsByTagName('head')[0].appendChild(style);
});