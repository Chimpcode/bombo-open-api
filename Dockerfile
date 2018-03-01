FROM golang:alpine
RUN apk add --no-cache git mercurial
ADD . /go/src/github.com/Chimpcode/bombo-open-api
ADD . /go/src/github.com/Chimpcode/bombo-open-api/files

RUN go get github.com/PuerkitoBio/goquery
RUN go get github.com/anaskhan96/soup
RUN go get github.com/gocolly/colly
RUN go get github.com/kataras/iris/...
RUN go get github.com/iris-contrib/middleware/cors

RUN go install github.com/Chimpcode/bombo-open-api

CMD ["/go/bin/bombo-open-api"]
EXPOSE 9800
