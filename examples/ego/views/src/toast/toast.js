var classActive = 'active';

var toast = function(msg, type, time) {
    if (type == undefined) {
        type = "error";
    };
    if (time == undefined) {
        time = 2000;
    };
    var toast = document.createElement('div');
    toast.classList.add('toast-container');
    var html = '<div class="toast-msg" id="toast-msg">' +
        '<i class="iconfont icon-' + type + '"></i>' +
        '<p class="toast-text">' + msg + '</p>' +
        '</div>';
    var docmsg = document.getElementById('toast-msg');
    if (docmsg == undefined) {
        toast.innerHTML = html;
    }
    toast.addEventListener('webkitTransitionEnd', function() {
        if (!toast.classList.contains(classActive)) {
            toast.parentNode.removeChild(toast);
        }
    });
    document.body.appendChild(toast);

    var mlength = msg.length;
    if (mlength > 5) {
        var mlen = (mlength - 5) / 1.5;
        var mrem = 5.5 + mlen;
        var docmsg = document.getElementById('toast-msg');
        docmsg.style.width = mrem + "rem";
    }
    toast.offsetHeight;
    toast.classList.add(classActive);
    setTimeout(function() {
        toast.classList.remove(classActive);
    }, time);
};
