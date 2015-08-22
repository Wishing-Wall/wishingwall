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
			<img class="img-responsive center-block" src="static/wishingwall.png" />
		</div>
	</div>
</div>

<div class="container-fluid">
	<div class="row-fluid">
			{{range $key, $val := .messages}}
				<div class="panel panel-primary">
					<div class="panel-heading ">{{$val.Id}}:</div>
					<div class="panel-body"><pre background="yellow">{{$val.Message}}</pre></div>
				</div>
			{{end}}
	</div>
</div>

<div class="container-fluid" id="PageButton">
	<div class="row-fluid">
		<div class="span12">
			<div class="btn-group">
				  <a href="?s={{.prestart}}&e={{.preend}}"><button class="btn btn-success" type="button">Prev<em class="icon-align-left"></em></button></a>
				 <a href="?s={{.nextstart}}&e={{.nextend}}"><button class="btn btn-success" type="button">Next<em class="icon-align-center"></em></button></a>
			</div>
		</div>
	</div>
</div>


<div class="container-fluid">
	<div class="row-fluid">
		<div class="span12">
			<form method="post">
				<fieldset>
					<legend>Share your hopes and dreams:</legend>
					<textarea type="text" name="clientmessage" cols=120 rows=10></textarea>
					<br>
					<button type="submit" class="btn btn-success">Submit</button>
				</fieldset>
			</form>
		</div>
	</div>
</div>


<div class="container-fluid" id="LG">
	<div class="row-fluid">
		<div class="span12">
			<p class="text-info text-center">
				<em>The Block Chain Height :</em> {{.maxblockindex}}
				<em>Wishing Wall Reached : </em> {{.parsedblockindex}}
			</p>
		</div>
	</div>
</div>

<div class="container-fluid" id="LG">
	<div class="row-fluid">
		<div class="span12">
			<p class="text-info text-center">
				<em>Warning </em> <p class="test-info text-center">This service operates automatically and can not delete your posts because the
     block chain cannot be rollback. The authors of this site are not responsible for 
	 the content it displays. Viewer discretion is advised.</p>

			</p>
		</div>
	</div>
</div>

</html>
