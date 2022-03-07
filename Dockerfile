FROM alpine:latest
LABEL maintainer="giuliohome@gmail.com"
COPY web .
COPY edit.html .
COPY view.html .
RUN chmod +x web
# https://stackoverflow.com/a/35613430
RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2 
EXPOSE 8080
CMD ["./web"]
# docker build --tag docker-golang-web .
# docker run -p:8080:8080  -it docker-golang-web
# open web preview on port 8080 https://8080-cs-162804068499-default.cs-europe-west1-iuzs.cloudshell.dev/edit/a1