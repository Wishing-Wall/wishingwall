<!DOCTYPE html>

<html>
<head>
<script>
  (function(i,s,o,g,r,a,m){i['GoogleAnalyticsObject']=r;i[r]=i[r]||function(){
  (i[r].q=i[r].q||[]).push(arguments)},i[r].l=1*new Date();a=s.createElement(o),
  m=s.getElementsByTagName(o)[0];a.async=1;a.src=g;m.parentNode.insertBefore(a,m)
  })(window,document,'script','//www.google-analytics.com/analytics.js','ga');

  ga('create', 'UA-67374398-1', 'auto');
  ga('send', 'pageview');

</script>

    <title>Wishing Wall</title>
    <link href="http://apps.bdimg.com/libs/bootstrap/3.3.0/css/bootstrap.min.css" rel="stylesheet">
    <script src="http://apps.bdimg.com/libs/jquery/2.0.0/jquery.min.js"></script>
    <script src="http://apps.bdimg.com/libs/bootstrap/3.3.0/js/bootstrap.min.js"></script>
	<script type="text/javascript">
        function $(id) {
            return document.getElementById(id);
        }
        function textToImg(txt) {
            var posy = 0;
			var posx = 0;
            
			var linespace = 8;
			var fontSize = 8;
            var fontWeight = 'normal';
            var canvas = $("canvas");

			eachlinetxt = txt.split('\r')
			console.log("eachlinetxt[0].length is ",eachlinetxt[0].length)
			console.log("eachlientxt.length is ", eachlinetxt.length)

            canvas.width = fontSize * eachlinetxt[0].length;
            canvas.height = linespace * eachlinetxt.length;
			console.log("canvas.width ", canvas.width)
			console.log("canvas.height ", canvas.height)
			
			
            var context = canvas.getContext('2d');
            context.clearRect(0, 0, canvas.width, canvas.height);
            context.font = fontSize  + 'px Lucida Console';
            //context.textBaseline = 'top';
            canvas.style.display = 'none';
            function fillTxt(text) {
		for (var i = 0; i < text.length; i++) {
                	context.fillText(text[i], i * fontSize, linespace * posy, canvas.width);
                }
		posy ++
	    }

            for ( var j = 0; j < eachlinetxt.length; j++) {
                fillTxt(eachlinetxt[j]);
            }
            var imageData = context.getImageData(0, 0, canvas.width, canvas.height);
			
			return canvas.toDataURL("image/png");
        }
	
		function TxtOrImg(val) {
			for (var j = 0; j < val.length; j++) {
				var body = $("txtorimg" + val[j].Id)
				if (val[j].BImg == 1)
				{
					body.innerHTML = "<div class=\"panel-body\"><pre><img src=\"" + textToImg(val[j].Message) + "\" class=\"img-responsive\" /></pre></div>"
				}else{
					body.innerHTML = "<div  class=\"panel-body\"><pre>" + val[j].Message + "</pre></div>"
				}
				
			}
		}
    </script>



</head>

<div class="container-fluid">
	<div class="row-fluid">
		<div class="span12">
			<img class="img-responsive center-block" src="static/wishingwall.png" />
		</div>
	</div>
</div>

<canvas id="canvas"></canvas>


<div class="container-fluid">
	<div class="row-fluid">
			{{range $key, $val := .messages}}
				<div class="panel panel-primary">
					<div class="panel-heading ">{{$val.Id}}:</div>
   					<div id="txtorimg{{$val.Id}}"></div>
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

	<script type="text/javascript">
		TxtOrImg({{.messages}})
    </script>


</html>
