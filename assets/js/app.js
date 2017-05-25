$(function(){
	// give thanks
	$(".btn-thx").click(function(){
		var dom_i = $(this).children("i:last-child");
			dom_i.text(Number(dom_i.text()) + 1);
		$(this).unbind("click");
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