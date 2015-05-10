mkdir gohttpd

cd gohttpd

export GOPATH=$PWD

go get github.com/nybuxtsui/gohttpd

cd bin

mkdir htdocs

echo "hello" > htdocs/index.html

./gohttpd
