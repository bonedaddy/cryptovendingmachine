pragma solidity ^0.5.10;

/// @title VendorManagement - Allows vendor product management
contract VendorManagement {

    enum State { NIL, REGISTERED }

    // Vendor allows vendor management
    struct Vendor {
        address owner;
        State state;
    }

    struct Product {
        bytes32 name;
        mapping(bytes32 => bool) soldAt;
    }

    struct Products {
        mapping(bytes32 => Product) index;
    }

    mapping(bytes32 => Vendor) public vendors;
    mapping(bytes32 => Products) internal products;

    event VendorRegistered(bytes32 _id, address _vendor);
    event ProductRegistered(bytes32 _name, bytes32[] _locations);

    function registerVendor() public returns (bool) {
        bytes32 id = getVendorID();
        require(notRegistered(id), "vendor must not be reigstered");
        vendors[id] = Vendor({
            owner: msg.sender,
            state: State.REGISTERED
        });
        emit VendorRegistered(id, msg.sender);
        return true;
    }

    function registerProduct(bytes32 _name, bytes32[] memory _locations) public returns (bool) {
        // get vendor id
        bytes32 id = getVendorID();
        // validate vendor is registered
        require(isRegistered(id), "vendor must be registered");
        // assign initial product name
        products[id].index[_name] = Product({name: _name});
        for (uint256 i = 0; i < _locations.length; i++) {
            products[id].index[_name].soldAt[_locations[i]] = true;
        }
        return true;
    }

    function addProductLocation(bytes32 _name, bytes32 _location) public returns (bool) {
        bytes32 id = getVendorID();
        require(isRegistered(id), "vendor must be registered");
        products[id].index[_name].soldAt[_location] = true;
        return true;
    }

    function removeProductLocation(bytes32 _name, bytes32 _location) public returns (bool) {
        bytes32 id = getVendorID();
        require(isRegistered(id), "vendor must be registered");
        products[id].index[_name].soldAt[_location] = false;
        return true;
    }

    // INTERNAL VIEW FUNCTIONS

    function notRegistered(bytes32 _id) internal view returns (bool) {
        if (vendors[_id].state == State.NIL) {
            return true;
        }
        return false;
    }

    function isRegistered(bytes32 _id) internal view returns (bool) {
        if (vendors[_id].state == State.REGISTERED) {
            return true;
        }
        return false;
    }

    function getVendorID() internal view returns (bytes32) {
        return keccak256(abi.encodePacked(msg.sender));
    }

}