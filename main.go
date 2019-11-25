package main

import (
	"os"
	"os/exec"
)

func main() {

	file, err := os.Create("/etc/nginx/conf.d/nginx.conf")

	if err != nil {
		panic(err)
	}

	config := `server {
 listen 80;
 listen [::]:80;

 server_name example.com;

 location /wrong/ {
    return 404;
 }
  
 location / {
   proxy_pass   https://strange-owl-31.localtunnel.me/;
 }
}`
	_, err = file.Write([]byte(config))

	if err != nil {
		panic(err)
	}

	err = file.Close()

	if err != nil {
		panic(err)
	}

	cmd := exec.Command("nginx", "-g", "daemon off;")
	err = cmd.Run()
}
