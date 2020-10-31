fetch('https://weibo.com/ajax/statuses/mymblog?uid=2821231635&page=1&feature=0').then(function(response) {
    return response.json()
}).then(function(data) {
    for (var item in data.data.list) {
        console.log(data.data.list[item].id)
        fetch('https://weibo.com/ajax/statuses/destroy',{
            method: 'POST',
            headers: {
                'content-type': 'application/json; charset=utf-8',
                'x-requested-with': 'XMLHttpRequest',
                'x-xsrf-token': 'I8WlwEXNyWAzBaRfRIzJGT0y'
            },
            body : '{"id" : "'+data.data.list[item].id+'"}'
        }).then(function(response) {
            return response.json()
        }).then(function(res){
            console.log(res.ok)
        })
    }
})