.PHONY: all clean

GOFLAGS=-gcflags "-N -l"

TARGET=crop

all:
	export CGO_ENABLED=0 && go build $(GOFLAGS) -o $(TARGET)

clean:
	-@rm -f gts
