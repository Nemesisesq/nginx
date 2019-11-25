
FROM golang:1.13 as build


WORKDIR /goroot
COPY main.go main.go

RUN go build -o main .

FROM nginx

RUN rm /etc/nginx/conf.d/default.conf
#COPY content /usr/share/nginx/html
#COPY conf/nginx.conf /etc/nginx/conf.d
#RUN cat /etc/nginx/conf.d/nginx.conf

COPY --from=build /goroot/main .

CMD ["./main"]







