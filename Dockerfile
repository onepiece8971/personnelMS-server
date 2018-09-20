FROM golang

MAINTAINER chenglinz <zhangchenglin@pxsj.com>

ENV PATH="${PATH}:/go/bin"

ARG PROXY="http://192.168.169.105:1080"
RUN git config --global http.proxy ${PROXY}

RUN go get -u -v github.com/kardianos/govendor \
	&& go get -v github.com/codegangsta/gin \
	&& go get -u -v github.com/derekparker/delve/cmd/dlv

WORKDIR ${GOPATH}/src/trunk/develop

# 2345 dlv
EXPOSE 3002 2345

# 开发模式
CMD ["gin", "-p", "3002", "&&", "gin", "run", "main.go"]
# CMD dlv --listen=:2345 --headless=true --api-version=2 exec ./xx