fibber: fibber.go size.go
	go build fibber.go size.go

size:
	cat size_preamble | sed 's/SIZE/${SIZE}/' > size.go

clean:
	rm -f fibber
