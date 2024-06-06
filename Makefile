.PHONY: all build run clean

all: build run clean

build:
	@go build -o=cfr cfr.go

run: build
	@./cfr

clean:
	@rm -f cfr
