FROM ubuntu
ADD . /tarserver
WORKDIR tarserver
EXPOSE 80
ENTRYPOINT ["bin/tarserver"]
