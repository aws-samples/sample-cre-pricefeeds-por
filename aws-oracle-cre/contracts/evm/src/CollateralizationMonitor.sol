// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "./keystone/IReceiver.sol";

contract CollateralizationMonitor is IReceiver {
    struct CollateralData {
        uint256 price;
        uint256 reserves;
        uint256 ratio;
        uint256 timestamp;
        bool isHealthy;
    }
    
    CollateralData public latestData;
    uint256 public minRatio = 120; // 120% = 1.2x (stored as percentage)
    
    event CollateralUpdated(
        uint256 price,
        uint256 reserves,
        uint256 ratio,
        bool isHealthy,
        uint256 timestamp
    );
    
    event ThresholdBreached(uint256 ratio, uint256 minRatio);
    
    function updateCollateral(
        uint256 _price,
        uint256 _reserves,
        uint256 _ratio,
        uint256 _timestamp,
        bool _isHealthy
    ) public {
        latestData = CollateralData({
            price: _price,
            reserves: _reserves,
            ratio: _ratio,
            timestamp: _timestamp,
            isHealthy: _isHealthy
        });
        
        emit CollateralUpdated(_price, _reserves, _ratio, _isHealthy, _timestamp);
        
        if (!_isHealthy) {
            emit ThresholdBreached(_ratio, minRatio);
        }
    }
    
    function getLatestData() external view returns (CollateralData memory) {
        return latestData;
    }
    
    function setMinRatio(uint256 _minRatio) external {
        minRatio = _minRatio;
    }
    
    // IReceiver interface implementation
    function onReport(bytes calldata metadata, bytes calldata report) external override {
        // The report contains the full encoded function call (selector + params)
        // Skip the first 4 bytes (function selector) and decode the parameters
        require(report.length >= 4, "Report too short");
        
        // Extract parameters starting from byte 4
        bytes calldata params = report[4:];
        (uint256 price, uint256 reserves, uint256 ratio, uint256 timestamp, bool isHealthy) = 
            abi.decode(params, (uint256, uint256, uint256, uint256, bool));
        updateCollateral(price, reserves, ratio, timestamp, isHealthy);
    }
    
    // IERC165 interface implementation
    function supportsInterface(bytes4 interfaceId) external pure override returns (bool) {
        return interfaceId == type(IReceiver).interfaceId || 
               interfaceId == type(IERC165).interfaceId;
    }
}
