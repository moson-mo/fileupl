package main

func documentation(url string) string {
	return `<html><body><span style="font-family:Courier New,Courier,monospace">
	<p><strong>fileupl</strong></p>
	
	Upload a file:<br/>
	<strong>$</strong> curl -H "password: your_pw" -F file=@file.name "` + url + `/upload"<br/><br/>
	
	Upload from stdin:<br/>
	<strong>$</strong> echo "something" | curl -H "password: your_pw" -F file=@- "` + url + `/upload"<br/><br/>
		
	It will the URL to the file, like:<br/>
	` + url + `/files/4e1243bd22c66e76c2ba9eddc1f91394e57f9f83</span></body>`
}

const uploadPage = `<!DOCTYPE html>
<html lang="en" data-bs-theme="dark">

<head>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <meta http-equiv="X-UA-Compatible" content="ie=edge" />
    <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.0-alpha1/dist/css/bootstrap.min.css" rel="stylesheet"
        integrity="sha384-GLhlTQ8iRABdZLl6O3oVMWSktQOp6b7In1Zl3/Jr59b6EGGoI1aFkw7cmDA6j6gD" crossorigin="anonymous">
    <title>fileupl</title>
</head>

<body>
    <div class="container" style="width: fit-content; margin-top: 40px;">
        <form enctype="multipart/form-data" action="/upload" method="post">
            <p class="d-flex justify-content-center"><strong>fileupl</strong> - File upload</p>

            <input class="form-control form-control-sm" type="file" name="file" /><br />
            <label class="form-label form-label-sm" for="password">Password:&nbsp;</label>
            <input class="form-control form-control-sm" id="password" type="password" name="password" /><br /><br />
            <input class="form-control form-control-sm" type="submit" value="Upload" />

        </form>
    </div>
    <script src="https://cdn.jsdelivr.net/npm/bootstrap@5.3.0-alpha1/dist/js/bootstrap.bundle.min.js"
        integrity="sha384-w76AqPfDkMBDXo30jS1Sgez6pr3x5MlQ1ZAGC+nuZB+EYdgRZgiwxhTBTkF7CXvN"
        crossorigin="anonymous"></script>

</body>

</html>`
