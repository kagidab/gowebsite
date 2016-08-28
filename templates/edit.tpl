<html>

<head>
	<meta http-equiv="content-type" content="text/html; charset=UTF-8">
	<title> Editing {{.Title}} </title>
	<link href="/css/wikistyle.css" type="text/css" rel="stylesheet">
	<script src="/js/jquery-2.1.4.min.js"></script>
	<script src="/js/wikiscripts.js"></script>
</head>

<body>

	<div id="sidebar">
		<div id="logo">
			<img src="/img/logo.png"></img>
		</div>
		<br> <br>
		<div id="sidetext">
			<a href="{{.WIKIPATH}}Main_page"> Main page </a> <br>
			<a href="{{.WIKIPATH}}List_all"> Page Listing </a> <br>
			<a href="{{.WIKIPATH}}Random_page"> Random </a> <br>
		</div>
	</div>
	<div id="tabs">
		<div id="left">
			<div class="off tab" id="article">
				<a href="{{.VIEWPATH}}{{.Title}}">Article</a>
			</div>
			<div class="on tab" id="edit">
				Edit
			</div>
		</div>
		
		<div id="right">
			<form action="javascript:redirect()">
				<input id="searchbar" placeholder="Search" />
			</form>
		</div>
	</div>
	<div id="span"></div>

	<div id="mainbar">
		<h1> Editing {{.Title}}<hr></h1>

		<form action='{{.SAVEPATH}}{{.Title}}' method="POST">
			<textarea id="write" name="body" rows="20" cols="80">{{printf "%s" .Body}}</textarea> 
			<div><input type="submit" value="Save"></div>
		</form>
	</div>

</body>
