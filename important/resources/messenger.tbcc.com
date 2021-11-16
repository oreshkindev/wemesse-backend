
server {  
	server_name messenger.tbcc.com;    
	return 301 https://messenger.tbcc.com$request_uri;
}

server {
	listen	443 ssl http2;
	
	server_name messenger.tbcc.com;
	
	ssl_certificate /etc/letsencrypt/live/messenger.tbcc.com/fullchain.pem;	
	ssl_certificate_key /etc/letsencrypt/live/messenger.tbcc.com/privkey.pem;
	ssl_trusted_certificate /etc/letsencrypt/live/messenger.tbcc.com/fullchain.pem;	

	include snippets/ssl-params.conf;    
	
	root /var/www/messenger.tbcc.com/html;
	index index.html;    

	location /index.html {
	try_files $uri $uri/ =404;
    }
	location / {
	proxy_pass http://localhost:9000;
    }
	location /release/ {
	try_files $uri $uri/ =404;
    }
	location /latest/ {
	try_files $uri $uri/ =404;
    }
}

