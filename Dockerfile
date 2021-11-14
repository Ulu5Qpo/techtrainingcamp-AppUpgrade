FROM centos:7
COPY techtrainingcamp-AppUpgrade /root/server
EXPOSE 9090
CMD /root/server
