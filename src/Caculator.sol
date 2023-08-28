// SPDX-License-Identifier: UNLICENSED
pragma solidity >=0.8.19;
// Example Solidity code with PUSH1 and ADD opcodes
contract SimpleCalculator {
    function add(uint256 a, uint256 b) public pure returns (uint256) {
        uint256 result;
        assembly {
            // Load the values of 'a' and 'b' onto the stack
            push1 a
            push1 b
            // Perform the addition operation
            add
            // Store the result
            result := mload(0x0)
        }
        return result;
    }
}
