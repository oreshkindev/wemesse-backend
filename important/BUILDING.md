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

#NOTE: Configuring authentication for SERVER mode.
#Enter the email address and password to use for the initial pgAdmin user account:
#Email address: user@domain.com
#Password:
#Retype password:
#Starting pgAdmin 4. Please navigate to http://127.0.0.1:5050 in browser.
#* Serving Flask app "pgadmin" (lazy loading)
#* Environment: production
#   WARNING: Do not use the development server in a production environment.
#   Use a production WSGI server instead.
#* Debug mode: off
```

## Installation guide on server (Ubuntu)

0. Create new linux user.
```shell
(root) groupadd wheel
(root) useradd wemesse -s /bin/bash -m -G wheel -c "TBCC Labs"
(root) passwd wemesse

# Add user to www-data group and to set rules to upload files
(root) usermod -a -G www-data wemesse
(root) mkdir -p /var/www/messenger.tbcc.com/html
(root) chown -R wemesse:www-data /var/www/messenger.tbcc.com/html
# Set rules to new user
(root) nano /etc/sudoers
(root) su - wemesse
mkdir /var/www/messenger.tbcc.com/html/source
```

1. Change permissions to main bin file
```shell
# Make executable bin
sudo chmod u+x /{path to bin}/main
```
2. Configure env.sh
```shell
#!/usr/bin/bash
#Server
export app_host=127.0.0.1
export app_port=9000
export dest_uri=http://182.92.107.179/wemesse/source/
export source_uri=https://messenger.tbcc.com/source/
export path_deploy=/var/www/messenger.tbcc.com/html/source

#Postgres
export user=postgres
export pass=postgres
export host=127.0.0.1
export port=5432
export name=tbcc_messenger
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

# === Some useful commands for PSQL ===
# Show all users
\du
# Show all the databases
\list or \l
# Show all of the psql tables
\d or \dt
# Choose database
\c tbcc_messenger
# Listing out tables using SELECT query
select * from table_name;
# Adding a record (INSERT INTO)
INSERT INTO table_name VALUES('1','7adfe73ef6a8744997bdec378ffaadcd');
# Delete one row from the table
DELETE FROM table_name WHERE id = 1;
# Delete table
DROP TABLE table_name;
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
# Add repository and install cetbot
sudo apt-get install python3-certbot-nginx
sudo apt install nginx
# Copy file 'messenger.tbcc.com' to /etc/nginx/sites-available/
sudo nano /etc/nginx/sites-available/messenger.tbcc.com
sudo ln -s /etc/nginx/sites-available/messenger.tbcc.com /etc/nginx/sites-enabled/
sudo nginx -t
sudo systemctl reload nginx
sudo systemctl enable nginx
sudo ufw enable
sudo ufw status
# Enable ufw OpenSSH port
sudo ufw allow '22/tcp'
# Enable ufw Zabbix port
sudo ufw allow '10050/tcp'
# Enable ufw Nginx port
sudo ufw allow 'Nginx Full'
# Obtain an SSL Certificate
sudo certbot --nginx -d messenger.tbcc.com
# Verify Certbot Auto-Renewal
sudo certbot renew --dry-run
```

# Errors while starting service

## Using newer libc on old Linux distributions
*/x86_64-linux-gnu/libc.so.6: version `GLIBC_2.32` not found*
```shell
# To check what version of glibc is installed use:
ldd --version
# Building glibc
mkdir $HOME/glibc/ && cd $HOME/glibc
wget http://ftp.gnu.org/gnu/libc/glibc-2.32.tar.gz
tar -xvzf glibc-2.32.tar.gz
mkdir build 
cd build
../configure --prefix=/opt/glibc-2.32
make -j4
sudo make install
# Now you should have glibc 2.32 installed in the installation dir. check with 
```
This will install glibc into */opt/glibc-2.32* but if run `ldd --version` it will still report the old version.
```shell
# Using the new glibc
LD_PRELOAD=/opt/glibc-2.32/lib/libc.so.6 ./main
# Syncing the glibc timezone

```

