$(function(){
	// give thanks
	$(".btn-thx").click(function(){
		var dom_i = $(this).children("i:last-child");
		var id = $(this).parent().attr("data-id");
			dom_i.text(Number(dom_i.text()) + 1);
			$(this).unbind("click");
		$.get("/article/add-thanks/" + id);
	});
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
})