### I'm building an EVM with ownership named OEVM.
1. For it, each opcode has its owner, which is represented by an address. During the running process of the smart contract (compiled), the EVN will track the usage of opcode of each owner like tracking gas consumption. And the tracking result will be recorded and used to calculate the contribution of each owner.

2. The owner for each opcode is the one who push the code into a repository(i, e. A solidity codebase).  So there maybe many contributors for one repository.

3. Could you help me to design OEVM as an expert of Golang, Solidity, EVM and Ethereum?

### Chatgpt

Designing a custom Ethereum Virtual Machine (EVM) with ownership tracking and opcode usage attribution involves multiple components, including the EVM implementation itself, a smart contract repository, and a mechanism to record and calculate opcode contributions. Here's a high-level design outline for your OEVM:

**1. Opcode Ownership and Tracking:**
You'll need to extend the existing EVM implementation to include opcode ownership and tracking. This might involve adding an ownership mapping for each opcode or opcode category. Additionally, you'll need to track opcode usage during contract execution. Consider creating a struct to store opcode ownership and usage data.

**2. Smart Contract Repository:**
Create a smart contract repository that allows contributors to push Solidity codebases. This repository should maintain a record of contributors and their contributions to each codebase. Contributors can be identified by their Ethereum addresses.

**3. Opcode Contribution Calculation:**
To calculate opcode contributions, follow these steps:

   - When a contributor pushes new code to the repository, the system should associate the contributor's address with the opcode ownership.
   - During contract execution on the OEVM, track opcode usage by incrementing the appropriate opcode counters in the ownership tracking data structure.
   - After contract execution, opcode contributions can be calculated based on the opcode counters and the proportion of opcode usage per contributor.

**4. Golang and Solidity Components:**

**Golang Components:**
- **EVM Core Modification:** Extend the existing EVM implementation (like go-ethereum) to include opcode ownership and tracking. Define the ownership tracking data structures and integrate opcode tracking logic into the execution flow.

- **Smart Contract Repository Backend:** Develop the backend to manage the smart contract repository. This includes functionalities to accept new codebase submissions, manage contributors, and maintain ownership associations.

**Solidity Components:**
- **Repository Smart Contract:** Develop a Solidity smart contract that allows contributors to submit codebases. This contract should manage ownership information and provide functions for calculating opcode contributions.

- **Contract Execution and Tracking:** Integrate the necessary hooks in the smart contract execution process to update opcode usage data during runtime.

**5. User Interface:**
Build a user interface that allows contributors to interact with the smart contract repository. Contributors can submit codebases and view their contributions. You can use web technologies for the front end, interacting with the smart contracts through Ethereum JSON-RPC APIs.

**6. Testing and Deployment:**
Thoroughly test your OEVM implementation and smart contracts on test networks. Ensure that opcode tracking, ownership attribution, and contribution calculations work as expected. Once testing is successful, deploy your system on the Ethereum mainnet or a chosen testnet.

Remember, this is a complex project that involves deep knowledge of Ethereum internals, Solidity, and system design. It's recommended to break down the project into smaller tasks and tackle them one by one. Additionally, consider collaborating with experts in Ethereum, EVM, Solidity, and system architecture to ensure the success of your OEVM project.
