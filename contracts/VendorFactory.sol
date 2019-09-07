pragma solidity ^0.5.10;

import "./VendorManagement.sol";

/// @title VendorFactory - Used to deploy vendor contracts
contract VendorFactory {

    // array of all vendor contracts registered
    address[] public vendorContracts;

    mapping(address => bool) public registeredVendors;
    // k1 = vendor manager address
    mapping(address => address) public vendorContract;

    function newVendor() public returns (bool) {
        // ensure they aren't registered
        require(!registeredVendors[msg.sender], "must not be registered");
        // deploy the vendor management contract
        VendorManagement contractAddr = new VendorManagement();
        // update vendor contract array
        vendorContracts.push(address(contractAddr));
        // update registered vendors
        registeredVendors[msg.sender] = true;
        // update vendor management address to contract mapping
        vendorContract[msg.sender] = address(contractAddr);
        return true;
    }
}