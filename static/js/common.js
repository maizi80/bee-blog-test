$(".btn-submit").on("click", function () {
    $(".btn-submit").attr("disabled", true);
    $.ajax({
        type: $('#post-form').attr('method'),
        url: $('#post-form').attr('action'),
        data: $('#post-form').serialize(),
        success: function (data){
            if (data.Code === 200){
                showTips('success')
                interval('success')
            } else {
                showTips('error')
                interval('error')
            }
            $(".alert").html(data.Msg)

        },
        error: function (res){
            $(".alert").html(res.statusText)
            showTips('error')
            interval('error')
        }
    })
    return false
})

function delRow(obj) {
    $(obj).attr("disabled", true).css("pointer-events","none");
    if (window.confirm("删除改数据后将不能恢复，是否继续？") === true){
        console.log($(obj).attr("url"))
        $.ajax({
            type: 'delete',
            url: $(obj).attr("url"),
            success: function (data){
                if (data.Code === 200){
                    showTips('success')
                    delInterval('success')
                } else {
                    showTips('error')
                    delInterval('error', obj)
                }
                $(".alert").html(data.Msg)

            },
            error: function (res){
                $(".alert").html(res.statusText)
                showTips('error')
                delInterval('error', obj)
            }
        })
    }
    return false
}

function showTips(type) {
    $(".alert").css("display", "block")
    if (type === 'success'){
        $(".alert").addClass("alert-success")
        $(".alert").removeClass("alert-danger")
    }
}

function interval(type) {
    let sec = 3
    if (type === 'success') sec = 1
    var interval = setInterval(function(){
        if (type === 'success') {
            window.history.go(-1);
        }
        $(".alert").css("display", "none")
        $(".btn-submit").attr("disabled", false);
        clearInterval(interval);
    }, sec*1000);
}

function delInterval(type, obj) {
    let sec = 3
    if (type === 'success') sec = 1
    var interval = setInterval(function(){
        if (type === 'success') {
            location.reload()
        }
        $(".alert").css("display", "none")
        $(obj).attr("disabled", false).css("pointer-events","none");
        clearInterval(interval);
    }, sec*1000);
}