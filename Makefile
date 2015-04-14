INSTALL_PATH=${HOME}/bin/

all: compile install

compile:
	mkdir -p dist
	go build -o dist/notify-serve notify-serve.go
	go build -o dist/notify notify.go

clean:
	rm -r dist/

install:
	cp dist/* ${INSTALL_PATH}
