.PHONY: vendor
vendor:
	@GO111MODULE=on go run -mod=vendor main.go --kubernetes-tag $(KUBERNETES_TAG) --kubernetes-path $(KUBERNETES_PATH)
