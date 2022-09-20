$(".btn-submit").on("click", function () {
    $(".btn-submit").attr("disabled", true);
    $.ajax({
        type: $('#post-form').attr('method'),
        url: $('#post-form').attr('action'),
        data: $('#post-form').serialize(),
        success: function (data){
            if (data.Code === 200){
                showTips('success')
            } else {
                showTips('error')
            }
            $(".alert").html(data.Msg)

        },
        error: function (res){
            $(".alert").html(res.statusText)
            showTips('error')
        }
    })
    return false
})

function showTips(type) {
    $(".alert").css("display", "block")
    if (type === 'success'){
        $(".alert").addClass("alert-success")
        $(".alert").removeClass("alert-danger")
        interval('success')
    } else {
        interval('error')
    }

}

function interval(type) {
    let sec = 3
    if (type === 'success') sec = 1
    var interval = setInterval(function(){
        if (type === 'success') {
            window.history.go(-1);
            // window.location.href = "/admin";
        }
        $(".alert").css("display", "none")
        $(".btn-submit").attr("disabled", false);
        clearInterval(interval);
    }, sec*1000);
}