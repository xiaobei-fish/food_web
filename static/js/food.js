$(document).ready(function () {
    //注册
    $("#register-from").validate({
        rules: {
            username: {
                required: true,
                rangelength: [5, 10]
            },
            password: {
                required: true,
                rangelength: [5, 10]
            },
            repassword: {
                required: true,
                rangelength: [5, 10],
                equalTo: "#register-password"
            }
        },
        messages: {
            username: {
                required: "请输入用户名",
                rangelength: "用户名必须是5-10位"
            },
            password: {
                required: "请输入密码",
                rangelength: "密码必须是5-10位"
            },
            repassword: {
                required: "请确认密码",
                rangelength: "密码必须是5-10位",
                equalTo: "两次输入的密码必须相等"
            },
            genre: {
                required: "请输入用户种类，0：买家，1：卖家",
                rangelength: "种类输入必须是1位(0或1)"
            }
        },
        submitHandler: function (form) {
            var urlStr = "/register";
            alert("urlStr:" + urlStr)
            $(form).ajaxSubmit({
                url: urlStr,
                type: "post",
                dataType: "json",
                success: function (data, status) {
                    alert("data:" + data.message + ":" + status)
                    if (data.code == 1) {
                        setTimeout(function () {
                            window.location.href = "/"
                        }, 1000)
                    }
                },
                error: function (data, status) {
                    alert("err:" + data.message + ":" + status)
                }
            })
        }
    });


    //登录
    $("#login-form").validate({
        rules: {
            username: {
                required: true,
                rangelength: [5, 10]
            },
            password: {
                required: true,
                rangelength: [5, 10]
            }
        },
        messages: {
            username: {
                required: "请输入用户名",
                rangelength: "用户名必须是5-10位"
            },
            password: {
                required: "请输入密码",
                rangelength: "密码必须是5-10位"
            }
        },
        submitHandler: function (form) {
            var urlStr = "/login"
            alert("urlStr:" + urlStr)
            $(form).ajaxSubmit({
                url: urlStr,
                type: "post",
                dataType: "json",
                success: function (data, status) {
                    alert("data:" + data.message + ":" + status)
                    if (data.code == 1) {
                        setTimeout(function () {
                            window.location.href = "/"
                        }, 1000)
                    }
                },
                error: function (data, status) {
                    alert("err:" + data.message + ":" + status)

                }
            });
        }
    });

    //修改和添加文章的表单
    $("#write-art-form").validate({
        rules: {
            store: "required",
            price: "required",
            intro: {
                required: true,
                minlength: 2
            }
        },
        messages: {
            store: "请输入商店名称",
            price: "请输入价格",
            intro: {
                required: "请输入简介",
                minlength: "简介内容最少两个字符"
            }
        },
        submitHandler: function (form) {
            var urlStr = "/food/add";
            //判断文章id确定提交的表单的服务器地址
            //若id大于零，说明是修改文章
            alert("urlStr:" + urlStr);
            $(form).ajaxSubmit({
                url: urlStr,
                type: "post",
                dataType: "json",
                success: function (data, status) {
                    alert("data:" + data.message);
                    setTimeout(function () {
                        window.location.href = "/food/add"
                    }, 1000)
                },
                error: function (data, status) {
                    alert("err:" + data.message + ":" + status)
                }
            });
        }
    });


    //文件
    $("#album-upload-button").click(function () {
        var filedata = $("#album-upload-file").val();
        if (filedata.length <= 0) {
            alert("请选择文件!");
            return
        }
        //文件上传通过Formdata去储存文件的数据
        var data = new FormData()
        data.append("upload", $("#album-upload-file")[0].files[0]);
        alert(data)
        var urlStr = "/upload"
        $.ajax({
            url: urlStr,
            type: "post",
            dataType: "json",
            contentType: false,
            data: data,
            processData: false,
            success: function (data, status) {
                alert("data:" + data.message);
            },
            error: function (data, status) {
                alert("err:" + data.message)
            }
        })
    });

    //修改密码表单
    $("#alter-password-form").validate({
        rules: {
            oldpassword: {
                required: true,
                rangelength: [5, 10]
            },
            newpassword: {
                required: true,
                rangelength: [5, 10]
            }
        },
        messages: {
            oldpassword: {
                required: "请输入旧密码",
                rangelength: "密码必须是5-10位"
            },
            newpassword: {
                required: "请确认新密码",
                rangelength: "密码必须是5-10位"
            },
        },
        submitHandler: function (form) {
            var urlStr = "/change";
            alert("urlStr:" + urlStr)
            $(form).ajaxSubmit({
                url: urlStr,
                type: "post",
                dataType: "json",
                success: function (data, status) {
                    alert("data:" + data.message + ":" + status)
                    if (data.code == 1) {
                        setTimeout(function () {
                            window.location.href = "/"
                        }, 1000)
                    }
                },
                error: function (data, status) {
                    alert("err:" + data.message + ":" + status)
                }
            })
        }
    });

    //修改信息表单
    $("#write-mes-form").validate({
        rules: {
            sex: "required",
            words: {
                required: true,
                minlength: 2
            }
        },
        messages: {
            sex: "请输入性别(男/女)",
            intro: {
                required: "请输入语录",
                minlength: "内容最少两个字符"
            }
        },
        submitHandler: function (form) {
            var urlStr = "/alter";
            alert("urlStr:" + urlStr);
            $(form).ajaxSubmit({
                url: urlStr,
                type: "post",
                dataType: "json",
                success: function (data, status) {
                    alert("data:" + data.message);
                    setTimeout(function () {
                        window.location.href = "/selfhome"
                    }, 1000)
                },
                error: function (data, status) {
                    alert("err:" + data.message + ":" + status)
                }
            });
        }
    });

    //头像文件
    $("#head-upload-button").click(function () {
        var filedata = $("#head-upload-file").val();
        if (filedata.length <= 0) {
            alert("请选择文件!");
            return
        }
        //文件上传通过Formdata去储存文件的数据
        var data = new FormData()
        data.append("head", $("#head-upload-file")[0].files[0]);
        alert(data)
        var urlStr = "/head"
        $.ajax({
            url: urlStr,
            type: "post",
            dataType: "json",
            contentType: false,
            data: data,
            processData: false,
            success: function (data, status) {
                alert("data:" + data.message);
            },
            error: function (data, status) {
                alert("err:" + data.message)
            }
        })
    });

    $("#add-form").validate({
        rules: {
            money: {
                required: true,
                rangelength: [2, 5]
            }
        },
        messages: {
            money: {
                required: "请输入充值金额",
                rangelength: "最少2位，最高5位"
            }
        },
        submitHandler: function (form) {
            var urlStr = "/admin/user";
            alert("urlStr:" + urlStr)
            $(form).ajaxSubmit({
                url: urlStr,
                type: "post",
                dataType: "json",
                success: function (data, status) {
                    alert("data:" + data.message + ":" + status)
                    if (data.code == 1) {
                        setTimeout(function () {
                            window.location.href = "/admin/user"
                        }, 1000)
                    }
                },
                error: function (data, status) {
                    alert("err:" + data.message + ":" + status)
                }
            })
        }
    });

    $("#comment-form").validate({
        rules: {
            comment: {
                required: true,
                rangelength: [2, 100]
            }
        },
        messages: {
            comment: {
                required: "请输入您的评论",
                rangelength: "最少2位，最多100位"
            }
        },
        submitHandler: function (form) {
            var comId = $("#comment-id").val();
            var urlStr = "/food/" + comId;
            alert("urlStr:" + urlStr)
            $(form).ajaxSubmit({
                url: urlStr,
                type: "post",
                dataType: "json",
                success: function (data, status) {
                    alert("data:" + data.message + ":" + status)
                    if (data.code == 1) {
                        setTimeout(function () {
                            location.reload()
                        }, 1000)
                    }
                },
                error: function (data, status) {
                    alert("err:" + data.message + ":" + status)
                }
            })
        }
    });


});
