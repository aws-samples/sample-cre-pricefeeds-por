// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "forge-std/Script.sol";
import "../src/CollateralizationMonitor.sol";

contract DeployCollateralizationMonitor is Script {
    function run() external {
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        
        vm.startBroadcast(deployerPrivateKey);
        
        CollateralizationMonitor monitor = new CollateralizationMonitor();
        
        console.log("CollateralizationMonitor deployed to:", address(monitor));
        
        vm.stopBroadcast();
    }
}
