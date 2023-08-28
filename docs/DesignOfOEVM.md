### I'm building an EVM with ownership named OEVM.
1. For it, each opcode has its owner, which is represented by an address. During the running process of the smart contract (compiled), the EVN will track the usage of opcode of each owner like tracking gas consumption. And the tracking result will be recorded and used to calculate the contribution of each owner.

2. The owner for each opcode is the one who push the code into a repository(i, e. A solidity codebase).  So there maybe many contributors for one repository.

3. Could you help me to design OEVM as an expert of Golang, Solidity, EVM and Ethereum?

### Chatgpt

