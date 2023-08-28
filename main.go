package main

import (
    "fmt"
    "strconv"
    "strings"
)

type Token struct {
    TokenType  string
    TokenValue string
}

func lexer(sourceCode string) []Token {
    // Split the source code into individual tokens
    tokens := strings.Split(sourceCode, " ")

    // Create a list of Token objects
    var tokenList []Token

    // Iterate over each token and classify its type
    for _, token := range tokens {
        if isNumeric(token) {
            tokenList = append(tokenList, Token{TokenType: "Numeric", TokenValue: token})
        } else {
            tokenList = append(tokenList, Token{TokenType: "Opcode", TokenValue: token})
        }
    }

    return tokenList
}

func isNumeric(token string) bool {
    _, err := strconv.Atoi(token)
    return err == nil
}

func analyzeSemantics(tokens []Token) {
    for i := 0; i < len(tokens); i++ {
        token := tokens[i]

        switch token.TokenValue {
        case "PUSH1":
            operand := tokens[i+1].TokenValue
            i++ // Skip the operand token
            analyzePUSH1(operand)

        case "ADD":
            analyzeADD()

        default:
            fmt.Println("Unknown token:", token.TokenValue)
        }
    }
}

func analyzePUSH1(operand string) {
    // Perform semantic analysis for PUSH1 opcode
    fmt.Println("Semantic analysis for PUSH1 opcode with operand:", operand)
    // Example semantic analysis logic for PUSH1
    // You can add your own custom logic here
    if isNumeric(operand) {
        fmt.Println("Operand is a valid numeric value")
    } else {
        fmt.Println("Operand is not a valid numeric value")
    }
}

func analyzeADD() {
    // Perform semantic analysis for ADD opcode
    fmt.Println("Semantic analysis for ADD opcode")
    // Example semantic analysis logic for ADD
    // You can add your own custom logic here
    fmt.Println("Performing addition operation")
}

func generateCode(tokens []Token) {
    for i := 0; i < len(tokens); i++ {
        token := tokens[i]

        switch token.TokenValue {
        case "PUSH1":
            operand := tokens[i+1].TokenValue
            i++ // Skip the operand token
            generatePUSH1(operand)

        case "ADD":
            generateADD()

        default:
            fmt.Println("Unknown token:", token.TokenValue)
        }
    }
}

func generatePUSH1(operand string) {
    // Generate code for PUSH1 opcode
    fmt.Println("Generated code for PUSH1 opcode with operand:", operand)
    // Example code generation logic for PUSH1
    // You can add your own custom logic here
    bytecode := "60" + operand // Add the PUSH1 opcode prefix (0x60) to the operand
    fmt.Println("Generated bytecode:", bytecode)
}

func generateADD() {
    // Generate code for ADD opcode
    fmt.Println("Generated code for ADD opcode")
    // Example code generation logic for ADD
    // You can add your own custom logic here
    bytecode := "01" // Placeholder bytecode for ADD operation
    fmt.Println("Generated bytecode:", bytecode)
}

func main() {
    sourceCode := `contract SimpleCalculator {
        function add(uint256 a, uint256 b) public pure returns (uint256) {
            uint256 result;
            assembly {
                // Load the values of 'a' and 'b' onto the stack
                PUSH1 a
                PUSH1 b
                // Perform the addition operation
                ADD
                // Store the result
                result := mload(0x0)
            }
            return result;
        }
    }`

    tokens := lexer(sourceCode)
    analyzeSemantics(tokens)
    generateCode(tokens)
}
