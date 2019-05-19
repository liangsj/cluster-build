From nginxconfd

RUN   mkdir -p /etc/confd/conf.d && mkdir -p /etc/confd/templates 

COPY ./conf/webapp_nginx.conf.toml /etc/confd/conf.d

COPY ./conf/8220.conf.tmpl  /etc/confd/templates 

COPY ./boost.sh /etc

run   chmod +x /etc/boost.sh 

