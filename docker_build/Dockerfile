FROM alpine
MAINTAINER X

ENV HOME_DIR /opt/aws_echo

WORKDIR ${HOME_DIR}

COPY ./aws_echo /opt/aws_echo

#RUN ln -s /etc/dgateway-dashboard/supervisord-dgateway-dashboard.conf /etc/supervisor/conf.d/supervisord-dgateway-dashboard.conf
#CMD supervisord -n -c /etc/supervisor/supervisord.conf

CMD ["/opt/aws_echo/aws_echo"]

