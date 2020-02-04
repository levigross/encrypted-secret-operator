.PHONY: generate

generate:
	operator-sdk generate crds
	operator-sdk generate openapi
