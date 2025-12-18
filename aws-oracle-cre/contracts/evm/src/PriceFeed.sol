// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "./keystone/IReceiver.sol";

contract PriceFeed is IReceiver {
    struct PriceData {
        uint256 price;
        uint256 timestamp;
    }
    
    PriceData public latestPrice;
    
    event PriceUpdated(uint256 price, uint256 timestamp);
    
    function updatePrice(uint256 _price, uint256 _timestamp) public {
        latestPrice = PriceData(_price, _timestamp);
        emit PriceUpdated(_price, _timestamp);
    }
    
    function getLatestPrice() external view returns (uint256, uint256) {
        return (latestPrice.price, latestPrice.timestamp);
    }
    
    // IReceiver interface implementation
    function onReport(bytes calldata metadata, bytes calldata report) external override {
        // The report contains the full encoded function call (selector + params)
        // Skip the first 4 bytes (function selector) and decode the parameters
        require(report.length >= 4, "Report too short");
        
        // Extract parameters starting from byte 4
        bytes calldata params = report[4:];
        (uint256 price, uint256 timestamp) = abi.decode(params, (uint256, uint256));
        updatePrice(price, timestamp);
    }
    
    // IERC165 interface implementation
    function supportsInterface(bytes4 interfaceId) external pure override returns (bool) {
        return interfaceId == type(IReceiver).interfaceId || 
               interfaceId == type(IERC165).interfaceId;
    }
}
