.PHONY: all
all: compile bindings

.PHONY: compile
compile:
	solc \
	--optimize \
	--optimize-runs 200 \
	--bin \
	--abi \
	--overwrite \
	-o bin/vendorManagement \
	contracts/VendorManagement.sol

.PHONY: bindings
bindings:
	abigen \
	--abi bin/vendorManagement/VendorManagement.abi \
	--bin bin/vendorManagement/VendorManagement.bin \
	--pkg vendormanagement \
	--out bindings/vendorManagement/bindings.go

# run standard go tooling for better readability
.PHONY: tidy
tidy: imports fmt
	go vet ./...
	golint ./...

# automatically add missing imports
.PHONY: imports
imports:
	find . -type f -name '*.go' -exec goimports -w {} \;

# format code and simplify if possible
.PHONY: fmt
fmt:
	find . -type f -name '*.go' -exec gofmt -s -w {} \;
