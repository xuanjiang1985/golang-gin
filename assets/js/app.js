$(function(){
	//ajax status hide
	$("#ajax-status").click(function(){
		$(this).hide();
	});
	//.disabled & add event function
	$(".disabled").click(function(){
		$("#ajax-status").text("已经没有了。");
		$("#ajax-status").show();
		$("#ajax-status").fadeOut(2000);
	});
	// comments button prev & next if have disabled
	function addDidsabledEve(ele) {
		$(ele).on('click',".disabled",function(){
			$("#ajax-status").text("已经没有了。");
			$("#ajax-status").show();
			$("#ajax-status").fadeOut(2000);
		});
	}

	// comments button prev & next if have no disabled and can run
	function scrollEve(ele) {
		$('html, body').animate({  
           	scrollTop: $(ele).offset().top - 100  
         },1000);
	}

	//create a short article
	$("#send").click(function(){
		var ctn = $("#content").val();
		var str = ctn.replace(/\ +/g,"");
		if(ctn == "" || str == ""){
			$("#content").val("").attr("placeholder","老大，我很饿....");
			return false;
		}
		$("#message-form").submit();
	});
	//get and display comments
	var done = [];
	$(".get-comments").click(function(){
		var comments = $(this).parent().next();
		if (comments.is(":hidden")) {
			comments.show();
		} else {
			comments.hide();
		}
		var article_id = $(this).parent().attr("data-id");
		if ($.inArray(article_id, done) == -1) {
			$.ajax({
				type:'get',
				url:'/article/get-comments/' + article_id,
				dataType:'json',
				success: function(data){
					//console.log(data);
					var content = "";
					$(data.comments).each(function(i,v){
						content += '<div><span class="author"><i class="fa fa-user" aria-hidden="true"></i> 匿名用户 &nbsp;&nbsp;<i>' + v.Created_at + '</i></span><p>' + v.Comment + '</p><hr></div>';
					});
	                $("#" + data.id).prev().html(content);
	                if(data.all_page == 1){
	                	$("#" + data.id).children("div").hide();
	                	return
	                }
	                if (data.current_page == 1) {
	                	$("#" + data.id).children("div").children("span").removeClass("disabled");
	                	$("#" + data.id).children("div").children(".comment-prev").addClass("disabled");
	                	addDidsabledEve($("#" + data.id).children("div"))
	                	return
	                }
	                if (data.current_page == data.all_page) {
	                	$("#" + data.id).children("div").children("span").removeClass("disabled");
	                	$("#" + data.id).children("div").children(".comment-next").addClass("disabled");
	                	addDidsabledEve($("#" + data.id).children("div"))
	                	return
	                }
	            },
	            error: function(data){
	                console.log(data);
	                return
	            }
			});
			done.push(article_id);
		}
	});
	//get next or prev page of comments
	$(".comment-next, .comment-prev").click(function(){
		if($(this).is(".disabled")){
			return
		}
		var article_id = $(this).parent().parent().attr("id");
		var current_page = $(this).parent().parent().data("currpage");
		if($(this).is(".comment-next")){
			var go_page = Number(current_page) + 1;
		} else {
			var go_page = Number(current_page) - 1;
		}
			$.ajax({
				type:'get',
				url:'/article/get-comments/' + article_id + '?page=' + go_page,
				dataType:'json',
				success: function(data){
					//console.log(data);
					var content = "";
					$(data.comments).each(function(i,v){
						content += '<div><span class="author"><i class="fa fa-user" aria-hidden="true"></i> 匿名用户 &nbsp;&nbsp;<i>' + v.Created_at + '</i></span><p>' + v.Comment + '</p><hr></div>';
					});
	                $("#" + data.id).prev().html(content);
	                if(data.all_page == 1){
	                	$("#" + data.id).children("div").hide();
	                	return
	                }
	                if (data.current_page == 1) {
	                	//$("#" + data.id).children("div").show();
	                	$("#" + data.id).children("div").children("span").removeClass("disabled");
	                	$("#" + data.id).children("div").children("a").text(data.current_page);
	                	$("#" + data.id).data("currpage",data.current_page);
	                	$("#" + data.id).children("div").children(".comment-prev").addClass("disabled");
	                	addDidsabledEve($("#" + data.id).children("div"));
	                	scrollEve($("#" + data.id).parent().parent());
	                	return
	                }
	                if (data.current_page >= data.all_page) {
	                	//$("#" + data.id).children("div").show();
	                	$("#" + data.id).children("div").children("span").removeClass("disabled");
	                	$("#" + data.id).children("div").children("a").text(data.current_page);
	                	$("#" + data.id).data("currpage",data.current_page);
	                	$("#" + data.id).children("div").children(".comment-next").addClass("disabled");
	                	addDidsabledEve($("#" + data.id).children("div"));
	                	scrollEve($("#" + data.id).parent().parent());
	                	return
	                }
	            },
	            error: function(data){
	                console.log(data);
	            }
			});
		});
	//post comments
	$(".btn-comment").click(function(){
		var comment = $(this).prev().val();
		var article_id = $(this).parent().attr("id");
		var str = comment.replace(/\ +/g,"");
		if(comment == "" || str == ""){
			$(this).prev().val("").attr("placeholder","请写下你的评论");
			return false;
		}
		$.ajax({
			type:"post",
			url: "/article/add-comment",
			headers: {"X-CSRF-TOKEN": $("input[name=_csrf]").val()},
			dataType:'json',
            data: {'comment':comment,'article_id':article_id},
            success: function(data){
                $("#" + data.id).children("input").val("");
                $("#" + data.id).prev().append('<div><span class="author"><i class="fa fa-user" aria-hidden="true"></i> 匿名用户 &nbsp;&nbsp;<i>' + data.created_at + '</i></span><p>' + data.comment + '</p><hr></div>');
                var ele_i = $("#" + data.id).parent().prev().find("i:last");
                ele_i.text(Number(ele_i.text()) + 1);
            },
            error: function(data){
                console.log(data);
            }
		});
	});
	//register validate
		$("#form-register").validate({
			rules: {
			    昵称: {
			    	required: true,
			    	minlength: 4
			    },
			    邮箱: {
			      required: true,
			      email: true
			    },
			    密码: {
			      required: true,
			      minlength: 6
			    },
			    密码确认: {
			    	required: true,
			      	minlength: 6,
			      	equalTo: "#re-password"
			    }
			  },
			  messages: {
			    昵称: {
			    	required: "不能为空",
			    	minlength: "至少4个字符"
			    },
			    邮箱: {
			      required: "不能为空",
			      email: "请填写正确的格式 如：name@domain.com"
			    },
			    密码: {
			      required: "不能为空",
			      minlength: "至少6位"
			    },
			    密码确认: {
			    	required: "不能为空",
			      	minlength: "至少6位",
			      	equalTo: "两次密码不相同"
			    }
			  }
		});
		//login image header
		$("#login").mouseenter(function(){
			$("#user-dropdown").show();
		}).mouseleave(function(){
			$("#user-dropdown").hide();
		});
		//post login
		$("#btn-login").click(function(){
			var email = $("#email").val();
			var password = $("#password").val();
			if (email == "") {
				$(this).prev().addClass("text-danger bg-danger").text("邮箱不能为空");
				return false;
			}
			var reg = /^[\w\-\.]+@[\w\-\.]+(\.\w+)+$/;
			    if (!reg.test(email)) {
			        $(this).prev().addClass("text-danger bg-danger").text("邮箱格式不对，如：name@domain.com");
			       return false;
			  }

			if (password == "") {
				$(this).prev().addClass("text-danger bg-danger").text("密码不能为空");
				return false;
			}

			if (password.length < 6) {
				$(this).prev().addClass("text-danger bg-danger").text("密码长度至少6位字符");
				return false;
			}

			$(this).prev().removeClass("text-danger bg-danger").html('<i class="fa fa-spinner fa-spin fa-2x" aria-hidden="true"></i> 登录中...');
			
			$.ajax({
			type:"post",
			url: "/login",
			headers: {"X-CSRF-TOKEN": $("input[name=_csrf]").val()},
			dataType:'json',
            data: {'email':email,'password':password},
            success: function(data){
            	if(data.error != ""){
            		$("#btn-login").prev().addClass("text-danger bg-danger").text(data.error);
            		return false;
            	}
      			window.location.reload();
            	},
            error: function(data){
                $("#btn-login").prev().addClass("text-danger bg-danger").text(data.statusText);
            	}
			});

		});
//change name
	$("#btn-set-name").click(function(){
		$.ajax({
			type:"post",
			url: "/auth/setting/name",
			headers: {"X-CSRF-TOKEN": $("input[name=_csrf]").val()},
			dataType:'json',
            data: {'name':$("input[name=name]").val()},
            success: function(data){
            	if(data.error != ""){
            		$("#ajax-status2").html('\
            		<div class="alert alert-danger">\
	            		<button type="button" class="close" data-dismiss="alert" aria-hidden="true">\
			                &times;\
			            </button>'
			             + data.error +
            		'</div>'
            		);
            	} else {
            		$("#ajax-status2").html('\
            		<div class="alert alert-success">\
	            		<button type="button" class="close" data-dismiss="alert" aria-hidden="true">\
			                &times;\
			            </button>'
			             + data.success +
            		'</div>'
            		);
            	}
            },
            error: function(data){
            	$("#ajax-status").text(data.statusText);
				$("#ajax-status").show();
				$("#ajax-status").fadeOut(2000);
            }
		});
	});
	//change sex setting
	$("#btn-set-sex").click(function(){
		$.ajax({
			type:"post",
			url: "/auth/setting/sex",
			headers: {"X-CSRF-TOKEN": $("input[name=_csrf]").val()},
			dataType:'json',
            data: {'sex':$("input[name=sex]:checked").val()},
            success: function(data){
            	if(data.error != ""){
            		$("#ajax-status2").html('\
            		<div class="alert alert-danger">\
	            		<button type="button" class="close" data-dismiss="alert" aria-hidden="true">\
			                &times;\
			            </button>'
			             + data.error +
            		'</div>'
            		);
            	} else {
            		$("#ajax-status2").html('\
            		<div class="alert alert-success">\
	            		<button type="button" class="close" data-dismiss="alert" aria-hidden="true">\
			                &times;\
			            </button>'
			             + data.success +
            		'</div>'
            		);
            	}
            },
            error: function(data){
            	$("#ajax-status").text(data.statusText);
				$("#ajax-status").show();
				$("#ajax-status").fadeOut(2000);
            }
		});
	});
	//upload header photo
	$("#btn-set-header").click(function(){
      $("#input-header").click();
    });
    $("#input-header").change(function(){
        /* 压缩图片 */
        lrz(this.files[0], {
            width: 44//设置压缩参数
        }).then(function (rst) {
            /* 处理成功后执行 */
            $("#ajax-status2").html('<i class="fa fa-spinner fa-spin fa-2x" aria-hidden="true"></i> 上传中...');
            rst.formData.append('base64img', rst.base64); // 添加额外参数
            $.ajax({
                url: "/auth/setting/header",
                type: "POST",
                data: rst.formData,
                headers: {"X-CSRF-TOKEN": $("input[name=_csrf]").val()},
                dataType:'json',
                processData: false,
                contentType: false,
                success: function (data) {
                	if (data.error != "") {
                		alert(data.error);
                		return
                	}
                    $("#img-header").attr("src", data.src);
                    $("#img-base-header").css("background-image",'url('+ data.src +')');
                    $("#ajax-status2").html('');
                },
                error: function(data){
                	alert(data.statusText);
                }
            });
        }).catch(function (err) {
            alert(err)
        }).always(function () {
            /* 必然执行 */
        })
    });
    //change password setting
    $("#btn-set-password").click(function(){
    	var oldpsd = $("input[name=old-psd]").val();
    	var newpsd = $("input[name=new-psd]").val();
    	var newpsd_confirm = $("input[name=new-psd-confirm]").val();
    	if (oldpsd == "" || newpsd == "" || newpsd_confirm =="") {
				$("#ajax-status2").html('\
            		<div class="alert alert-danger">\
	            		<button type="button" class="close" data-dismiss="alert" aria-hidden="true">\
			                &times;\
			            </button>\
			             密码不能为空\
            		</div>'
            		);
				return false;
			}
			if (oldpsd.length < 6 || newpsd.length < 6 || newpsd_confirm < 6) {
				$("#ajax-status2").html('\
            		<div class="alert alert-danger">\
	            		<button type="button" class="close" data-dismiss="alert" aria-hidden="true">\
			                &times;\
			            </button>\
			             长度至少6位\
            		</div>'
            		);
				return false;
			}

			if (newpsd != newpsd_confirm) {
				$("#ajax-status2").html('\
            		<div class="alert alert-danger">\
	            		<button type="button" class="close" data-dismiss="alert" aria-hidden="true">\
			                &times;\
			            </button>\
			             两次新密码不相同\
            		</div>'
            		);
				return false;
			}

			$.ajax({
			type:"post",
			url: "/auth/setting/password",
			headers: {"X-CSRF-TOKEN": $("input[name=_csrf]").val()},
			dataType:'json',
            data: {'oldpsd':oldpsd,'newpsd':newpsd,'newpsd_confirm':newpsd_confirm},
            success: function(data){
            	if(data.error != ""){
            		$("#ajax-status2").html('\
            		<div class="alert alert-danger">\
	            		<button type="button" class="close" data-dismiss="alert" aria-hidden="true">\
			                &times;\
			            </button>'
			             + data.error +
            		'</div>'
            		);
            	} else {
            		$("#ajax-status2").html('\
            		<div class="alert alert-success">\
	            		<button type="button" class="close" data-dismiss="alert" aria-hidden="true">\
			                &times;\
			            </button>'
			             + data.success +
            		'</div>'
            		);
            	}
            },
            error: function(data){
            	$("#ajax-status").text(data.statusText);
				$("#ajax-status").show();
				$("#ajax-status").fadeOut(2000);
            }
		});
    });
});

// give thanks
	$(".btn-thx").click(function(){
		var dom_i = $(this).children("i:last-child");
		var id = $(this).parent().attr("data-id");
			dom_i.text(Number(dom_i.text()) + 1);
			$(this).unbind("click");
		$.get("/article/add-thank/" + id);
	});
	$("#img-base-header").click(function(){
		if($("#user-dropdown").is(":hidden")){
			$("#user-dropdown").show();
		} else {
			$("#user-dropdown").hide();
		}
		
	});



