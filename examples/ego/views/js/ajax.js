function ajax(options) {
    options = options || {};
    options.type = (options.type || "GET").toUpperCase();
    options.dataType = options.dataType || "json";
    var params = formatParams(options.data);

    if (window.XMLHttpRequest) {
        var xhr = new XMLHttpRequest();
    } else {
        var xhr = new ActiveXObject('Microsoft.XMLHTTP');
    }

    xhr.onreadystatechange = function() {
        if (xhr.readyState == 4) {
            var status = xhr.status;
            if (status >= 200 && status < 300) {
                options.success && options.success(xhr.responseText, xhr.responseXML);
            } else {
                options.fail && options.fail(status);
            }
        }
    }

    if (options.type == "GET") {
        xhr.open("GET", options.url + "?" + params, true);
        xhr.send(null);
    } else if (options.type == "POST") {
        xhr.open("POST", options.url, true);

        xhr.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");
        xhr.send(params);
    }
}

function formatParams(data) {
    var arr = [];
    for (var name in data) {
        arr.push(encodeURIComponent(name) + "=" + encodeURIComponent(data[name]));
    }
    arr.push(("v=" + Math.random()).replace(".", ""));
    return arr.join("&");
}

function jsonp(options) {
    options = options || {};
    if (!options.url || !options.callback) {
        throw new Error("Parameter is not legal !");
    }

    var callbackName = ('jsonp_' + Math.random()).replace(".", "");
    var oHead = document.getElementsByTagName('head')[0];
    options.data[options.callback] = callbackName;
    var params = formatParams(options.data);
    var oscr = document.createElement('script');
    oHead.appendChild(oscr);

    window[callbackName] = function(json) {
        oHead.removeChild(oscr);
        clearTimeout(oscr.timer);
        window[callbackName] = null;
        options.success && options.success(json);
    };

    oscr.src = options.url + '?' + params;

    if (options.time) {
        oscr.timer = setTimeout(function() {
            window[callbackName] = null;
            oHead.removeChild(oscr);
            options.fail && options.fail({ message: "Time out !" });
        }, time);
    }
};