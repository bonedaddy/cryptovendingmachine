pragma solidity ^0.5.10;
pragma experimental ABIEncoderV2;

interface VendorManagementI {
  function products (string calldata) external view returns (string memory name, uint256 cost);
  function withdrawFunds () external returns (bool);
  function soldAt (string calldata, string calldata) external view returns (bool);
  function withdrawLock () external view returns (bool);
  function removeProductLocation (string calldata _name, string calldata _location) external returns (bool);
  function addProductLocation (string calldata _name, string calldata _location) external returns (bool);
  function owner () external view returns (address);
  function id () external view returns (bytes32);
  function registerProduct (string calldata _name, string[] calldata _locations, uint256 _cost) external returns (bool);
}
