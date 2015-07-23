FROM golang:1.4.2
MAINTAINER Ivan Kurakin <ivancurachin@gmail.com>
RUN go get github.com/constabulary/gb/...
RUN mkdir -p /srv/services/apartment_search
WORKDIR /srv/services/apartment_search
CMD ["go-wrapper", "run"]
COPY . /srv/services/apartment_search
RUN gb build all
ENTRYPOINT /srv/services/apartment_search/bin/apartment_search