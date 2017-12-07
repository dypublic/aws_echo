FROM phusion/baseimage:0.9.22
MAINTAINER X

#RUN    apt-get update                                                                                                  \
#    && apt-get install -y wget                                                                                         \
#    && apt-get install -y git                                                                                          \
#    && apt-get install -y ca-certificates                                                                              \
#    && apt-get install -y apt-transport-https                                                                          \
#    && apt-get autoclean                                                                                               \
#    && apt-get --purge -y autoremove
RUN    apt-get update                                              \
    && apt-get install -y supervisor                               \
    && apt-get install -y python                                   \
    && rm -rf /tmp/* /var/tmp/*                                    \
    && rm -rf /var/lib/apt/lists/*                                 \
    && rm -f /etc/ssh/ssh_host_*                                   \
    && apt-get autoclean && apt-get --purge autoremove	

RUN mkdir -p /var/log/supervisor /etc/supervisor/conf.d
COPY ./supervisord.conf /etc/supervisor/supervisord.conf

#end
