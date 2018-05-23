pragma solidity ^0.4.0;

contract HelloWorldContract {
	uint result;
	function sayHi() constant returns (uint) {
		for(uint i = 0; i < 100; i++) {
			result = result +1;
		}
		return result;
	}
}
