#!/usr/bin/bash
#Server FOR PRODUCTION
export app_host=127.0.0.1
export app_port=9000
#Do not use / in the end of path !!!!
#Exp http://182.92.107.179/wemesse/source
export dest_uri=http://182.92.107.179/wemesse/source
#Exp http://localhost:9000/source
export source_uri=http://localhost:9000/source
#Exp /var/www/messenger.tbcc.com/html/source
export path_deploy=/var/www/messenger.tbcc.com/html/source

#Postgres FOR PRODUCTION
export user=postgres
export pass=postgres
export host=127.0.0.1
export port=5432
export name=tbcc_messenger