// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "forge-std/Script.sol";
import "../src/PriceFeed.sol";

contract DeployPriceFeed is Script {
    function run() external {
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        
        vm.startBroadcast(deployerPrivateKey);
        
        PriceFeed priceFeed = new PriceFeed();
        
        console.log("PriceFeed deployed to:", address(priceFeed));
        
        vm.stopBroadcast();
    }
}
