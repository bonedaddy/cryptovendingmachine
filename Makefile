.PHONY: compile
compile:
	solc --optimize --optimize-runs 200 --bin --abi contracts/VendorManagement.sol --overwrite -o bin/vendorManagement