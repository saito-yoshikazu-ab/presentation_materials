.PHONY:lint
lint:
	go fmt ./result/...
	goimports -local app -w ./result
	golangci-lint run --timeout=10m ./result