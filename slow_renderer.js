var userAgent = window.navigator.userAgent.toLowerCase();
var ios = /iphone|ipod|ipad/.test(userAgent),
    android = /android/.test(userAgent),
    mobile = /mobile/.test(userAgent),
    gecko = /gecko\//.test(userAgent);

if( (!gecko) && (ios || android || mobile)) {
    document.body.id="mobile-webkit";
}