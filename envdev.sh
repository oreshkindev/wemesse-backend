#!/usr/bin/bash
#Server ONLY FOR DEVELOPMENT MODE
export app_host=127.0.0.1
export app_port=9000
#Do not use / in the end of path !!!!
#Exp http://182.92.107.179/wemesse/source
export dest_uri=http://182.92.107.179/wemesse/source
#Exp https://messenger.tbcc.com/source
export source_uri=https://messenger.tbcc.com/source
#Exp /var/www/messenger.tbcc.com/html/source
export path_deploy=source

#Postgres ONLY FOR DEVELOPMENT MODE
export user=oresh
export pass=root
export host=127.0.0.1
export port=5432
export name=postgres