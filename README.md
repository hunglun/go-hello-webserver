# Source 
https://blog.devgenius.io/go-docker-hello-world-f092ecf7cead

# Cross compile
Go program can be easily cross compiled to different target.
See Makefile

For example, to make the server for Q8, run 
`make arm64`

# Using React Web App
Copy the build directory to client subfolder.

For example,
```
git clone git@github.com:SpeQtral/device-control-webapp.git 
cd device-control-webapp
npm run build
cp -a build <go-server-root-dir>/client/build
```
Now you can run the go server, serving the web app!
`go run react_app_webserver.go`

To deploy it to Q8,
cross compile the webserver, by using Makefile.
Then create the following folder structure in target device
```
- arm64_webserver
- client
    - build
```
# Using App Web Server with environment variables
 PORT=8000 REACT_APP_BUILD_DIR=~/device_control_webapp/client/build go run react_app_web_server.go 


# TODO 
add a server handler for device access from webapp.
