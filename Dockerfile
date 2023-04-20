FROM ubuntu:20.04

RUN \
    apt update && \
    apt -y upgrade && \
    apt install build-essential wget sudo curl make -y && \
    rm -rf /usr/local/go && tar -C /usr/local -xzf go1.19.4.linux-amd64.tar.gz && \

ENV PATH $PATH:/usr/local/go/bin
# export PATH=$PATH:/usr/local/go/bin

WORKDIR /home/

# RUN apt install elixir -y