$('.table tr').each(function(index){
    var row = $(this).find("td").last().children().attr('href');
    var str = String(row)
    if (str == 'undefined') {
        return
    }
    var last =str.lastIndexOf('/')
    var url = 'http://auth.security.xxxx.com/workflow/audit'
    console.log(str)
    var data = '{"taskId": "'+str.substring(last+1)+'", "bizId": "'+str.substring(16,22)+'", "action": "pass", "comment": "同意"}'
    console.log($.parseJSON(data))
    $.ajax({
        url:url,
        type:"POST",
        data:data,
        contentType:"application/json; charset=utf-8",
        dataType:"json",
        success: function(data){
            console.log('success')
        },
        error: function (jqXHR, textStatus, errorThrown){
            console.log(errorThrown)
        }
    })
})


