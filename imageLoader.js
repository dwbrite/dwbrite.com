
var srcs;
initImages();

function initImages()
{
  var nscrpt = document.getElementsByClassName("media-noscript");
  var length = nscrpt.length;
  srcs = new Array(length);
  for (i = 0; i < length; i++) {
  
    var btn = document.createElement("button");
    btn.innerHTML = "load media";
    btn.setAttribute("onclick", "showbtn("+i+");");
    btn.setAttribute("id", "media-btn-"+i);
    btn.setAttribute("class", "media-btn");
    btn.setAttribute("type", "button");
    btn.setAttribute("title", "Load Content");
  
    var div = document.createElement("div");
    div.setAttribute("class", "media-div");
  
    var img = document.createElement("img");
    img.setAttribute("id", "media-img-"+i);
    img.setAttribute("class", "media-img");

    var a = document.createElement("a");
    a.setAttribute("id", "media-anchor-"+i);
    a.setAttribute("class", "media-anchor");
    a.appendChild(img);
  
    div.appendChild(btn);
    div.appendChild(a);
  
    srcs[i] = nscrpt[i].getAttribute("id");
  	nscrpt[i].parentElement.insertBefore(div, nscrpt[i]);
  }
}

function showbtn(i) {
    var button = document.getElementById("media-btn-"+i);
    button.setAttribute("onclick", "hidebtn("+i+");");
    button.innerHTML = "hide media";
    button.setAttribute("title", "Hide Content");

    var a = document.getElementById("media-anchor-"+i);
    a.setAttribute("href", srcs[i]);
    a.setAttribute("style", "display:block");

    var img = document.getElementById("media-img-"+i);
    img.setAttribute("src", srcs[i]);
}

function hidebtn(i) {
    var button = document.getElementById("media-btn-"+i);
    button.setAttribute("onclick", "showbtn("+i+");");
    button.innerHTML = "view media";
    button.setAttribute("title", "View Content");

    var a = document.getElementById("media-anchor-"+i);
    //img.removeAttribute("src");
    a.setAttribute("style", "display:none");

    //var img = document.getElementById("media-img-"+i);
    //img.removeAttribute("src");
    //img.setAttribute("style", "display:none");
}
