APP_NAME := adventofcode2025

.PHONY: build run clean

build:
	go build -o bin/$(APP_NAME) .

run:
	@if [ -z "$(CHALLENGE)" ] || [ -z "$(PART)" ] || [ -z "$(INPUT)" ]; then \
		echo "Usage: make run CHALLENGE=<number> PART=<1|2> INPUT=<path-to-input>"; \
		exit 1; \
	fi
	go run . -challenge $(CHALLENGE) -part $(PART) -input $(INPUT)

clean:
	rm -rf bin
