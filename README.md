### fileupl

A simple http service allowing you to upload and share files.

Some use cases:
- Share files within you local network
- Install it on your server and self-host your own pastebin service.
- Quickly share screenshots, logfiles, etc. with others.

The `/upload` endpoint can be protected by an `APIKey` and thus require authentication.  
`SHA-1` hash of the file content is used as the new filename, accessible via the `/files` endpoint.  
If you want to preserve the original name of the file, add a query param: `?pfn=true` to the URL.

Place a `index.html` into the directory where you store the files to disable content listing.

#### Configuration

A configuration example can be found in the assets directory.  
By default, fileupl expects a config file at `/etc/fileupl/config.toml`.

You can specify an alterative location of you config file with the `-c` parameter.

```toml
# URL that is being used
# We need this for the documentation and return the URL to an uploaded file
url = "http://127.0.0.1:9999"

# Interface/port that is being listened on
# :port will listen on all network interfaces
listen_address = ":9999"

# API key for authentication
# Leaving it blank will disable auth
api_key = "changeme"

# Path to the directory where files are being stored
directory = "/tmp/fileupl"

# Maximum filesize for uploaded files in MB
max_mb = 10

```

#### Usage

The simplest way is to use curl for uploading files.  
Take a look at the [examples](EXAMPLES.md) and the client script [upl](assets/upl)
