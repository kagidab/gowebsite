function redirect(){
  var regex = /^[a-zA-Z \_]+$/;
  console.log(regex.exec($("#searchbar").val()));
	if(regex.exec($("#searchbar").val())){
		window.location.href = "/wiki/" + $("#searchbar").val();
	}
}
