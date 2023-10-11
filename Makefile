export REMOTE_IP="35.91.228.41"

run:
	go run main.go

status:
	@curl -s localhost:8080/status | jq

# URL=xyz make encode
encode:
	@curl -s -X POST -d '{"url": "${URL}"}' http://localhost/encode | jq

remote-encode:
	@curl -s -X POST -d '{"url": "${URL}"}' http://${REMOTE_IP}/encode | jq

scp:
	scp -r static ubuntu@${REMOTE_IP}:~
	scp main ubuntu@${REMOTE_IP}:~

amd64:
	GOOS=linux GOARCH=amd64 go build -o main .

m1:
	GOOS=darwin GOARCH=arm go build .
