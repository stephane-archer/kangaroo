kangaroo: main.go
	go build

kangarooMacArm64: export CGO_ENABLED = 1
kangarooMacArm64: export GOOS = darwin
kangarooMacArm64: export GOARCH = arm64
kangarooMacArm64: main.go
	go build -o $@ main.go

.PHONY: clean
clean:
	go clean
