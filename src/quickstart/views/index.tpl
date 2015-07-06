<!DOCTYPE html>

<html>
<head>
  <title>Wishing Wall</title>
</head>
<body>
  <head>
    <pre>
                         _       __ _        __     _                  _       __        __ __
                        | |     / /(_)_____ / /_   (_)____   ____ _   | |     / /____ _ / // /
                        | | /| / // // ___// __ \ / // __ \ / __ `/   | | /| / // __ `// // /
                        | |/ |/ // /(__  )/ / / // // / / // /_/ /    | |/ |/ // /_/ // // /
                        |__/|__//_//____//_/ /_//_//_/ /_/ \__, /     |__/|__/ \__,_//_//_/
                                                          /____/

    </pre>
  </head>
  <body>
	{{range $key, $val := .messages}}
		<div style="font-family:verdana">{{$val.Id}}:<pre>{{$val.Message}}</pre></div>
	{{end}}
  </body>
  <br>
  <tr>
      <td><a href="?s={{.prestart}}&e={{.preend}}">Prev</a></td>
	  <td><a href="?s={{.nextstart}}&e={{.nextend}}">Next</a></td>
  </tr>

  <br>
  <form method="post">
	  Share your hopes and dreams:<br>
	  <textarea type="text" name="clientmessage" cols=120 rows=10></textarea>
	  <br>
	  <input type="submit" value="submit">
  <br>
  <br>
  <br>
  <br>
  <br>

  <footer>This service operates automatically and can not delete your posts because the
     block chain cannot be rollback. The authors of this site are not responsible for 
	 the content it displays. Viewer discretion is advised.</footer>
</html>
