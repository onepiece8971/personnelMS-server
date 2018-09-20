FROM golang

MAINTAINER chenglinz <zhangchenglin@pxsj.com>

ENV PATH="${PATH}:/go/bin"

RUN go get -u -v github.com/kardianos/govendor \
	&& go get -v github.com/codegangsta/gin

WORKDIR ${GOPATH}/src

EXPOSE 3000

CMD ["gin", "run", "main.go"]