## Work enviropment for developing on host machine (Arch Linux)

*Install Postman*
```shell
# Install snap from the Arch User Repository
git clone https://aur.archlinux.org/snapd.git
cd snapd
makepkg -si
# enable the systemd unit
sudo systemctl enable --now snapd.socket
# enable the classic snap support
sudo ln -s /var/lib/snapd/snap /snap
# install Postman
snap install postman
# download postman desktop agent https://www.postman.com/downloads/
tar -xzf Postman-linux-x64-7.32.0.tar.gz
sudo mkdir -p /opt/postman
sudo mv Postman /opt/postman
sudo ln -s /opt/postman/Postman /usr/local/bin/postman
postman

# create a desktop file
sudo nano ~/.local/share/applications/Postman.desktop

# Enter the following content in the file—replacing opt
[Desktop Entry]
Encoding=UTF-8
Name=Postman
Exec=/opt/postman/Postman %U
Icon=/opt/postman/app/resources/app/assets/icon.png
Terminal=false
Type=Application
Categories=Development;
```

*Install PostgreSQL*
```shell
sudo pacman -Sy
sudo pacman -S postgresql
# PostgreSQL is not running
sudo systemctl status postgresql
# login as the postgres
sudo su - postgres
# initialize the data directory
initdb --locale en_US.UTF-8 -D /var/lib/postgres/data
exit
sudo systemctl start postgresql
sudo systemctl enable postgresql
# PostgreSQL now is running
sudo systemctl status postgresql
```

*Install pgadmin4*
```shell
sudo mkdir /var/lib/pgadmin
sudo mkdir /var/log/pgadmin
sudo chown $USER /var/lib/pgadmin
sudo chown $USER /var/log/pgadmin
python3 -m venv pgadmin4
source pgadmin4/bin/activate
(pgadmin4) pip install pgadmin4
(pgadmin4) sudo pgadmin4
```
NOTE: Configuring authentication for SERVER mode.
Enter the email address and password to use for the initial pgAdmin user account:
Email address: user@domain.com
Password:
Retype password:
Starting pgAdmin 4. Please navigate to http://127.0.0.1:5050 in browser.
* Serving Flask app "pgadmin" (lazy loading)
* Environment: production
   WARNING: Do not use the development server in a production environment.
   Use a production WSGI server instead.
* Debug mode: off

## Installation guide on server (Ubuntu)

0. Create new linux user.
```shell
groupadd wheel
useradd wemesse -s /bin/bash -m -G wheel -c "TBCC Labs"
passwd wemesse

# Set rules to new user
nano /etc/sudoers
```

1. Change permissions to main bin file
```shell
# make executable bin
sudo chmod u+x /{path to bin}/main
```
2. Configure env.yaml
```shell
# Server configuration
host: 127.0.0.1
port: 9000
uri: https://messenger.tbcc.com/release/
tmp: ./tmp/
release: /var/www/messenger.tbcc.com/html/release/
notes: ./release-notes/
duration: 30 # Time in seconds
salt: 8781e03169eed720a768ce7eecfc6a21

# Database configurations
database:
  user: postgres
  pass: postgres
  host: 127.0.0.1
  port: 5432
  table: tbcc_messenger

# POST Downgrade 
# curl -X POST curl -X POST http://45.77.55.28:9000/api/v1/downgrade/8781e03169eed720a768ce7eecfc6a21/[version]
# https://messenger.tbcc.com/api/v1/updates/2264_00
```

3. Create new database.
*Install postgres*
```shell
# Ubuntu install
sudo apt install postgresql postgresql-contrib
# Enter to postgres role
sudo -u postgres psql
# Change password
ALTER USER postgres PASSWORD 'postgres';
# Create db
CREATE DATABASE tbcc_messenger;
```

4. Configure reload service.
*Create wemesse service*
```shell
# Create /ect/systemd/system/wemesse.service
# and paste following code:

[Unit]
Description=WeMesse by TBCC backend service
After=network.target

[Service]
Type=simple
WorkingDirectory=/opt/weMesse
ExecStart=/opt/weMesse/main
Restart=on-failure

[Install]
WantedBy=multi-user.target
```

5. Star and enable link to service
*Start wemesse service*
```shell
sudo systemctl start wemesse.service
sudo systemctl enable wemesse.service
sudo systemctl status wemesse.service
```

6. Build certbot, lets'ecrypt and  nginx services.
*Start nginx service*
```shell
# Create folder of web:
sudo mkdir -p /var/www/messenger.tbcc.com/html/release/ 
# Add repository and install cetbot
sudo add-apt-repository ppa:certbot/certbot
sudo apt install python-certbot-nginx
# Copy file 'messenger.tbcc.com' to /etc/nginx/sites-available/
sudo nano /etc/nginx/sites-available/messenger.tbcc.com
sudo nginx -t
sudo systemctl reload nginx
# Enable ufw Nginx port
sudo ufw status
sudo ufw allow 'Nginx Full'
# Obtain an SSL Certificate
sudo certbot --nginx -d messenger.tbcc.com
# Verify Certbot Auto-Renewal
sudo certbot renew --dry-run
```




