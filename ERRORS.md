# MYSQL Connection failed due to native authentication not enabled

# Set Native authentication 
**Error: golang**
starting access to database from go lang
[mysql] 2021/09/02 17:23:26 connector.go:95: could not use requested auth plugin 'mysql_native_password': this user requires mysql native password authentication.
2021/09/02 17:23:26 this user requires mysql native password authentication.
exit status 1

        AllowNativePasswords: true, # <- This 

