FROM debian:latest

COPY ./file.sh /
WORKDIR /

RUN apt-get update
RUN apt-get install htop
# # Since htop is a terminal application we also have to set xterm. See links for more info:
# # https://stackoverflow.com/questions/27826241/running-nano-in-docker-container
# # https://github.com/moby/moby/issues/9299
# # https://en.wikipedia.org/wiki/Xterm
ENV TERM xterm 
ENTRYPOINT ["/file.sh"]