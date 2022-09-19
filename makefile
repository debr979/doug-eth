GO ?=go
BUILD =$(GO) build
TIDY = $(GO) mod tidy

init:
	rm main log.out

tidy:
	$(TIDY)

dev: tidy;
	$(GO) run main.go

build: tidy;
	$(BUILD) -o main main.go