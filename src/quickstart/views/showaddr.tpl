<!DOCTYPE html>

<html>
<head>
    <title>Wishing Wall</title>
    <link href="http://apps.bdimg.com/libs/bootstrap/3.3.0/css/bootstrap.min.css" rel="stylesheet">
    <script src="http://apps.bdimg.com/libs/jquery/2.0.0/jquery.min.js"></script>
    <script src="http://apps.bdimg.com/libs/bootstrap/3.3.0/js/bootstrap.min.js"></script>
</head>


<div class="container-fluid">
	<div class="row-fluid">
		<div class="span12">
			<div class="hero-unit ">
				<h1  class="text-center">
					Please send at least
				</h1>
				<h1 class="text-center">
					{{.minmoney}}
				</h1>
				<h1 class="text-center">
					bitcoins to address
				</h1>
				<h1 class="text-center">
					{{.address}}
				</h1>
				<h1 class="text-center">
					Your message will be "written" on Blockchain after 2 confirms
				</h1>
			</div>
		</div>
	</div>
</div>

</html>