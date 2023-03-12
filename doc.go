package main

func documentation(url string) string {
	return `<html><body><span style="font-family:Courier New,Courier,monospace">
	<p><strong>fileupl</strong></p>
	
	Upload a file:<br/>
	<strong>$</strong> curl -H "APIKey: your_key" -F file=@file.name "` + url + `/upload"<br/><br/>
	
	Upload from stdin:<br/>
	<strong>$</strong> echo "something" | curl -H "APIKey: your_key" -F file=@- "` + url + `/upload"<br/><br/>
		
	It will the URL to the file, like:<br/>
	` + url + `/files/4e1243bd22c66e76c2ba9eddc1f91394e57f9f83</span></body>`
}
