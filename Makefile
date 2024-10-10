install:
	go build -o bin/bodo main.go
	mkdir -p /etc/bodo
	mkdir -p /usr/local/share/bodo
	cp index.html themes.yaml /usr/local/share/bodo/
	cp -n config.yaml /etc/bodo
	cp bin/bodo /usr/local/bin/bodo

uninstall:
	rm -r /usr/local/share/bodo
	rm /usr/local/bin/bodo
