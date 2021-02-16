FROM golang:latest

LABEL maintainer Lenin <lenin.ribeiroquadros@gmail.com>

RUN mkdir /build
WORKDIR /build

RUN export GO111MODULE=on
RUN go get github.com/leninribeiro/mconf-api-lenin
RUN cd /build && git clone https://github.com/leninribeiro/mconf-api-lenin.git

RUN cd /build/mconf-api-lenin && go build

ENV PORT 8000

EXPOSE 8000

ENTRYPOINT [ "/build/mconf-api-lenin/mconf-api-lenin" ]