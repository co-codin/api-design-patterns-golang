run:
	go run ./cmd/web

test:
	cd cmd/web && go test -v -count=1 .
