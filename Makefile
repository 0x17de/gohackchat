.PHONY: all clean
SRCS := $(wildcard cmd/gohackbot/*.go)
EXTRA_SRCS := $(wildcard pkg/hack/*.go)
EXTRA_SRCS += $(wildcard cmd/gohackbot/commands/*.go)

all: gohackbot

clean:
	rm gohackbot

gohackbot: $(SRCS) $(EXTRA_SRCS)
	go build -o $@ $(SRCS)
