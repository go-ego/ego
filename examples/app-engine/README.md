# Guide to run ego under App Engine LOCAL Development Server

1. Download, install and setup Go in your computer. (That includes setting your `$GOPATH`.)
2. Download SDK for your platform from here: `https://developers.google.com/appengine/downloads?hl=es#Google_App_Engine_SDK_for_Go`
3. Download ego source code using: `$ go get github.com/go-ego/ego`
4. Navigate to examples folder: `$ cd $GOPATH/src/github.com/go-ego/ego/examples/`
5. Run it: `$ goapp serve app-engine/`