package constants

const (
	INDEX_HTML_PATH = "./internal/html/index.html"
	ROOT_HTML_PATH  = "./internal/html/root.html"
)

// These are not currently in use...

const HTML_SEGMENT = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <script src="https://unpkg.com/htmx.org@1.9.11" integrity="sha384-0gxUXCCR8yv9FM2b+U3FDbsKthCI66oH5IA9fHppQq9DDMHuMauqq1ZHBpJxQ0J0" crossorigin="anonymous"></script>
    <title>Webhook Server</title>
</head>
<body style="background-color:#222222; padding:5px 30px 5px 30px;">
    <h1 style="color:azure;">Triggered webhooks</h1>
    <div id="webhooks-container" style="display:flex; justify-content:start; flex-wrap:wrap; row-gap:20px; column-gap:30px;">
        {{ range .Webhooks }}
        
        {{ end }}
    </div
</body>
</html>`

const POST_HTML_SEGMENT = `<div style="width:27%; height:auto; background-color:#2f4f4f; padding:10px 30px 10px 30px; border-radius:10px;">
    <h2 style="text-align:center; background-color:#ac329c; color:azure; padding:5px; border-radius:10px;">{{ .Method }}</h2>
    <p style="color:azure;">{{ .Url }}</p>
    <p style="color:azure;">{{ .Timestamp }}</p>
    <p style="color:azure;">{{ .Content }}</p>
</div>`

const GET_HTML_SEGMENT = `<div style="width:27%; height:auto; background-color:#2f4f4f; padding:10px 30px 10px 30px; border-radius:10px;">
    <h2 style="text-align:center; background-color:#53b037; color:azure; padding:5px; border-radius:10px;">{{ .Method }}</h2>
    <p style="color:azure;">{{ .Url }}</p>
    <p style="color:azure;">{{ .Timestamp }}</p>
    <p style="color:azure;">{{ .Content }}</p>
</div>`

const DELETE_HTML_SEGMENT = `<div style="width:27%; height:auto; background-color:#2f4f4f; padding:10px 30px 10px 30px; border-radius:10px;">
    <h2 style="text-align:center; background-color:#cb2525; color:azure; padding:5px; border-radius:10px;">{{ .Method }}</h2>
    <p style="color:azure;">{{ .Url }}</p>
    <p style="color:azure;">{{ .Timestamp }}</p>
    <p style="color:azure;">{{ .Content }}</p>
</div>`

const PATCH_PUT_HTML_SEGMENT = `<div style="width:27%; height:auto; background-color:#2f4f4f; padding:10px 30px 10px 30px; border-radius:10px;">
    <h2 style="text-align:center; background-color:#f0e320; color:azure; padding:5px; border-radius:10px;">{{ .Method }}</h2>
    <p style="color:azure;">{{ .Url }}</p>
    <p style="color:azure;">{{ .Timestamp }}</p>
    <p style="color:azure;">{{ .Content }}</p>
</div>`
