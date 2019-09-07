pragma solidity ^0.5.10;

/// @title VendorManagement - Allows vendor product management
contract VendorManagement {
    address public owner;
    bytes32 public id;

    struct Product {
        bytes32 name;
        uint256 cost;
    }

    mapping(bytes32 => Product) public products;
    // k1 = product
    // k2 = vending machine
    mapping(bytes32 => mapping(bytes32 => bool)) public soldAt;

    event ProductRegistered(bytes32 _name, bytes32[] _locations);


    constructor() public {
        id = keccak256(abi.encodePacked(msg.sender));
    }

    function registerProduct(bytes32 _name, bytes32[] memory _locations, uint256 _cost) public returns (bool) {
        require(onlyVendor(), "caller must be vendored");
        // assign initial product name
        products[_name] = Product({name: _name, cost: _cost});
        for (uint256 i = 0; i < _locations.length; i++) {
            soldAt[_name][_locations[i]] = true;
        }
        return true;
    }

    function addProductLocation(bytes32 _name, bytes32 _location) public returns (bool) {
        require(onlyVendor(), "caller must be vendored");
        soldAt[_name][_location] = true;
        return true;
    }

    function removeProductLocation(bytes32 _name, bytes32 _location) public returns (bool) {
        require(onlyVendor(), "caller must be vendored");
        soldAt[_name][_location] = true;
        return true;
    }

    // INTERNAL VIEW FUNCTIONS
    function onlyVendor() internal view returns (bool) {
        if (msg.sender == owner) {
            return true;
        }
        return false;
    }
}