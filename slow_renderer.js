var userAgent = window.navigator.userAgent.toLowerCase();
var ios = /iphone|ipod|ipad/.test(userAgent),
    android = /android/.test(userAgent),
    mobile = /mobile/.test(userAgent),
    gecko = /gecko\//.test(userAgent);

if( (!gecko) && (ios || android || mobile)) {
    document.body.style.background = "linear-gradient(-45deg, rgba(24, 0, 128, .5), rgba(255, 128, 0, .6))";
}