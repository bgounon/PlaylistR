FROM golang:1.12

WORKDIR /go/src/PlaylistR
COPY src/* ./

RUN go get -v

CMD PlaylistR -id $PLAYLIST_ID -api $API_KEY
