pragma solidity ^0.5.10;

import "./SafeMath.sol";
import "./Interfaces/VendorManagementI.sol";

/// @title VendingMachine - Represents a single, physical vending machine
contract VendingMachine {
    using SafeMath for uint256;

    // vendors selling items through the machine
    mapping(address => bool) public vendors;
    mapping(address => address) public vendorContracts;

    function addVendor(address _vendorContract) public returns (bool) {
        require(!vendors[msg.sender], "vendor must not be registered");
        VendorManagementI vmI = VendorManagementI(_vendorContract);
        require(vmI.owner() == msg.sender, "caller must be vendor manager");
        vendors[msg.sender] = true;
        vendorContracts[msg.sender] = _vendorContract;
        return true;
    }

}