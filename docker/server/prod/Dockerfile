FROM golang:1.17.0 as builder

RUN apt-get update && apt-get install -y sudo
# amazon-ssm-agent インストール
RUN mkdir /tmp/ssm
RUN wget https://s3.amazonaws.com/ec2-downloads-windows/SSMAgent/latest/debian_amd64/amazon-ssm-agent.deb -O /tmp/ssm/amazon-ssm-agent.deb \
  && sudo dpkg -i /tmp/ssm/amazon-ssm-agent.deb \
  && cp /etc/amazon/ssm/seelog.xml.template /etc/amazon/ssm/seelog.xml
COPY docker-entrypoint.sh /
ENTRYPOINT [ "/docker-entrypoint.sh" ]