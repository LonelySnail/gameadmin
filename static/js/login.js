
var hosts = "http://192.168.1.225:8000";
/*
 * 手机号正则表达式
 */
function isPhoneNo(phone) {
    var pattern = /^1[34578]\d{9}$/;
    return pattern.test(phone);
}

/*
 * 邮箱正则表达式
 */
function isEmailNo(email) {
    var pattern = /^[a-z0-9]+([._\\-]*[a-z0-9])*@([a-z0-9]+[-a-z0-9]*[a-z0-9]+.){1,63}[a-z0-9]+$/;
    return pattern.test(email);
}

/*
 * 验证手机号
 */
function checkPhone() {
    if ($.trim($('#telephone').val()).length == 0) {
        alert('请输入手机号');
        $('#telephone').focus();
        return false;
    } else {
        if (isPhoneNo($('#telephone').val()) == false) {
            alert('手机号码不正确');
            $('#telephone').focus();
            return false;
        }
    }
    return true;
}

/*
 * 验证密码
 */
function checkPassword() {
    if ($.trim($('#password').val()).length == 0) {
        alert('没有输入密码');
        $('#password').focus();
        return false;
    } else {
        var rePass = $.trim($('#rePassword').val());
        var pass = $.trim($('#password').val());
        if (rePass.length != 0) {
            if (rePass != pass) {
                alert('两次密码不一致');
                return false;
            }
        }
    }
    return true;
}

/*
 * 重复密码
 */
function checkrePassword() {
    var rePass = $.trim($('#rePassword').val());
    if (rePass.length == 0) {
        alert('密码没有输入');
        $('#rePassword').focus();
        return false;
    } else {
        var pass = $.trim($('#password').val());
        if (pass.length == 0) {
            $('#rePassword').val('');
            alert('密码没有输入');
            $('#password').focus();
            return false;
        } else if (rePass != pass) {
            alert('两次密码不一致');
            return false;
        }
    }

    return true;

}

/*
 * 验证联系人
 */
function checkName() {
    if ($.trim($('#username').val()).length == 0) {
        alert('请输入账户');
        $('#username').focus();
        return false;
    }

    return true;

}


/*
 * 注册表单提交时验证
 */
function DoRegister() {
    if (checkPhone() == false) {
        return false;
    }

    if (checkPassword() == false) {
        return false;
    }

    if (checkName() == false) {
        return false;
    }


    data = {"name":$('#username').val(),"password":$('#password').val(),"rePassword":$('#rePassword').val(),"phone":$('#telephone').val()};
    $.ajax({
        type: 'POST',
        url: hosts+"/doRegister",
        data: data,
        success: function (obj) {
            if (obj.toString() == "success") {
                if (confirm("注册成功 请登录")) {
                    $(window).attr('location',hosts+"/");
                }
            }else {
                alert(obj.toString());
            }

        },
        error: function (e) {
            alert(e.toString());
        }

    });


}

function register() {
    window.open(hosts+"/register");
}

function logout() {
    $(window).attr('location',hosts+"/");
}