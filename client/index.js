var url = "http://127.0.0.1:9001/users";


data = {
    nickname: "243affsfdsssaf22",
    email:"testdfsd4325@gmail.com",
    password: "tesfsaft2"
};

jQuery.ajax({
    type: "POST",
    url: url,
    data: JSON.stringify(data),
    contentType:"application/json; charset=utf-8",
    //dataType: "text/plain; charset=UTF-8",
    success: console.log(JSON.stringify(data))

});

$.postJSON = function(url, data, callback) {
    return jQuery.ajax({
        'type': 'POST',
        'url': url,
        'contentType': 'application/json; charset=utf-8',
        'data': JSON.stringify(data),
        'dataType': '',
        'success': console.log(data)
    });
};


