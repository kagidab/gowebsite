function selectNick(){
	nickname = $("#nick").val();
	$("#nicky").hide();
	$("#chat").show();

	(function(){
		$.get("read", 
				function (data) {
					var obj = $("#chatbox").text(data)
					$("#chatbox").html(obj.html().replace(/\n/g,'<br/>')); 
				}, "html")
		setTimeout(arguments.callee, 1000)
	})();
}

function write(){
	$.get("write", {say : nickname + ": " + $("#say").val()}) 
		$("#say").val("");
}

window.onload = function(){
	$("#selectnick").click( function(){ selectNick() });
	$("#sendbutton").click( function(){ write() });
}
