# cfcomputing

This is it...the most flexible web application.

## Built with :heart: and help from the most awesome world of open source

Thank you! https://github.com/

Thank you! https://stackoverflow.com/

Thank you! https://letsencrypt.org/

```
Donating to ISRG / Let's Encrypt:   https://letsencrypt.org/donate
Donating to EFF:                    https://eff.org/donate-le
```

Thank you! https://www.nginx.com/

Thank you! https://nodejs.org/

Thank you! https://vuejs.org/

Thank you! https://postgresql.org/

Thank you to every developer of every open source package/module/gist/snippet/....etc....whatever you want to call it....thank you for your help!

## General Notes for myself

Install nginx on centos 7
https://www.digitalocean.com/community/tutorials/how-to-install-nginx-on-centos-7

```
sudo yum install epel-release
sudo yum install nginx
sudo systemctl start nginx

sudo firewall-cmd --permanent --zone=public --add-service=http
sudo firewall-cmd --permanent --zone=public --add-service=https
sudo firewall-cmd --reload
```

Installing a cert from letsencrypt
https://certbot.eff.org/lets-encrypt/centosrhel7-nginx

For wildcard need this but didn't do it at this time

`sudo certbot -a dns-plugin -i nginx -d "*.example.com" -d example.com --server https://acme-v02.api.letsencrypt.org/directory`
