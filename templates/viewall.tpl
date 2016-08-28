<html>

<head>
	<meta http-equiv="content-type" content="text/html; charset=UTF-8">
	<title> {{.Title}} </title>
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
			<div class="on tab" id="article">
				Article Listing
			</div>
		</div>
		
		<div id="right">
			<form id="search" action="javascript:redirect()">
				<input id="searchbar" placeholder="Search" />
			</form>
		</div>
	</div>
	<div id="span"></div>

	<div id="mainbar">
		<h1> {{.Title}}<hr></h1>
			{{range .Links}}<a href="{{ . }}">{{ . }}</a><br>{{end}}
		</form>
	</div>

</body>

