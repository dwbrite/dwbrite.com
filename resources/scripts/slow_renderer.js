var userAgent = window.navigator.userAgent.toLowerCase();
var gecko = /gecko\//.test(userAgent);

if(!gecko) {
    document.body.id="mobile-webkit";
}