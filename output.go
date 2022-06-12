package solc

import (
	"encoding/json"
)

type Output struct {
	Errors    []Error                        `json:"errors,omitempty"`
	Sources   map[string]SourceOut           `json:"sources,omitempty"`
	Contracts map[string]map[string]Contract `json:"contracts,omitempty"`
}

type Error struct {
	SourceLocation   SourceLocation `json:"sourceLocation,omitempty"`
	Type             string         `json:"type,omitempty"`
	Component        string         `json:"component,omitempty"`
	Severity         string         `json:"severity,omitempty"`
	Message          string         `json:"message,omitempty"`
	FormattedMessage string         `json:"formattedMessage,omitempty"`
}

type SourceLocation struct {
	File  string `json:"file,omitempty"`
	Start int    `json:"start,omitempty"`
	End   int    `json:"end,omitempty"`
}

type SourceOut struct {
	ID        int             `json:"id,omitempty"`
	AST       ASTsolc         `json:"ast,omitempty"`
	LegacyAST json.RawMessage `json:"legacyAST,omitempty"`
}

type Contract struct {
	ABI      []json.RawMessage `json:"abi,omitempty"`
	Metadata string            `json:"metadata,omitempty"`
	UserDoc  json.RawMessage   `json:"userdoc,omitempty"`
	DevDoc   json.RawMessage   `json:"devdoc,omitempty"`
	IR       string            `json:"ir,omitempty"`
	// StorageLayout StorageLayout     `json:"storageLayout,omitempty"`
	EVM   EVM   `json:"evm,omitempty"`
	EWASM EWASM `json:"ewasm,omitempty"`
}

type EVM struct {
	Assembly          string                       `json:"assembly,omitempty"`
	LegacyAssembly    json.RawMessage              `json:"legacyAssembly,omitempty"`
	Bytecode          Bytecode                     `json:"bytecode,omitempty"`
	DeployedBytecode  Bytecode                     `json:"deployedBytecode,omitempty"`
	MethodIdentifiers map[string]string            `json:"methodIdentifiers,omitempty"`
	GasEstimates      map[string]map[string]string `json:"gasEstimates,omitempty"`
}

type Bytecode struct {
	Object         string                                `json:"object,omitempty"`
	Opcodes        string                                `json:"opcodes,omitempty"`
	SourceMap      string                                `json:"sourceMap,omitempty"`
	LinkReferences map[string]map[string][]LinkReference `json:"linkReferences,omitempty"`
}

type LinkReference struct {
	Start int `json:"start,omitempty"`
	End   int `json:"end,omitempty"`
}

type EWASM struct {
	Wast string `json:"wast,omitempty"`
	Wasm string `json:"wasm,omitempty"`
}

type ASTsolc struct {
	AbsolutePath    string `json:"absolutePath"`
	ExportedSymbols struct {
		One interface{} `json:"One"`
	} `json:"exportedSymbols"`
	ID       int64  `json:"id"`
	NodeType string `json:"nodeType"`
	Nodes    []struct {
		Abstract                bool          `json:"abstract"`
		BaseContracts           []interface{} `json:"baseContracts"`
		ContractDependencies    []interface{} `json:"contractDependencies"`
		ContractKind            string        `json:"contractKind"`
		Documentation           interface{}   `json:"documentation"`
		FullyImplemented        bool          `json:"fullyImplemented"`
		ID                      int64         `json:"id"`
		LinearizedBaseContracts []int64       `json:"linearizedBaseContracts"`
		Literals                []string      `json:"literals"`
		Name                    string        `json:"name"`
		NodeType                string        `json:"nodeType"`
		Nodes                   []struct {
			Body struct {
				ID         int64  `json:"id"`
				NodeType   string `json:"nodeType"`
				Src        string `json:"src"`
				Statements []struct {
					Expression struct {
						ArgumentTypes    interface{} `json:"argumentTypes"`
						HexValue         string      `json:"hexValue"`
						ID               int64       `json:"id"`
						IsConstant       bool        `json:"isConstant"`
						IsLValue         bool        `json:"isLValue"`
						IsPure           bool        `json:"isPure"`
						Kind             string      `json:"kind"`
						LValueRequested  bool        `json:"lValueRequested"`
						NodeType         string      `json:"nodeType"`
						Src              string      `json:"src"`
						Subdenomination  interface{} `json:"subdenomination"`
						TypeDescriptions struct {
							TypeIdentifier string `json:"typeIdentifier"`
							TypeString     string `json:"typeString"`
						} `json:"typeDescriptions"`
						Value string `json:"value"`
					} `json:"expression"`
					FunctionReturnParameters int64  `json:"functionReturnParameters"`
					ID                       int64  `json:"id"`
					NodeType                 string `json:"nodeType"`
					Src                      string `json:"src"`
				} `json:"statements"`
			} `json:"body"`
			Documentation    interface{} `json:"documentation"`
			FunctionSelector string      `json:"functionSelector"`
			ID               int64       `json:"id"`
			Implemented      bool        `json:"implemented"`
			Kind             string      `json:"kind"`
			Modifiers        []struct {
				ID           int64 `json:"id"`
				ModifierName struct {
					ID                    int64  `json:"id"`
					Name                  string `json:"name"`
					NodeType              string `json:"nodeType"`
					ReferencedDeclaration int64  `json:"referencedDeclaration"`
					Src                   string `json:"src"`
				} `json:"modifierName"`
				NodeType string `json:"nodeType"`
				Src      string `json:"src"`
			} `json:"modifiers"`
			Name       string      `json:"name"`
			NodeType   string      `json:"nodeType"`
			Overrides  interface{} `json:"overrides"`
			Parameters struct {
				ID         int64  `json:"id"`
				NodeType   string `json:"nodeType"`
				Parameters []struct {
					Constant         bool   `json:"constant"`
					ID               int64  `json:"id"`
					Mutability       string `json:"mutability"`
					Name             string `json:"name"`
					NameLocation     string `json:"nameLocation"`
					NodeType         string `json:"nodeType"`
					Scope            int64  `json:"scope"`
					Src              string `json:"src"`
					StateVariable    bool   `json:"stateVariable"`
					StorageLocation  string `json:"storageLocation"`
					TypeDescriptions struct {
						TypeIdentifier string `json:"typeIdentifier"`
						TypeString     string `json:"typeString"`
					} `json:"typeDescriptions"`
					TypeName struct {
						ID               int64  `json:"id"`
						Name             string `json:"name"`
						NodeType         string `json:"nodeType"`
						Src              string `json:"src"`
						StateMutability  string `json:"stateMutability"`
						TypeDescriptions struct {
							TypeIdentifier string `json:"typeIdentifier"`
							TypeString     string `json:"typeString"`
						} `json:"typeDescriptions"`
					} `json:"typeName"`
					Visibility string `json:"visibility"`
				} `json:"parameters"`
				Src string `json:"src"`
			} `json:"parameters"`
			ReturnParameters struct {
				ID         int64  `json:"id"`
				NodeType   string `json:"nodeType"`
				Parameters []struct {
					Constant         bool   `json:"constant"`
					ID               int64  `json:"id"`
					Mutability       string `json:"mutability"`
					Name             string `json:"name"`
					NodeType         string `json:"nodeType"`
					Scope            int64  `json:"scope"`
					Src              string `json:"src"`
					StateVariable    bool   `json:"stateVariable"`
					StorageLocation  string `json:"storageLocation"`
					TypeDescriptions struct {
						TypeIdentifier string `json:"typeIdentifier"`
						TypeString     string `json:"typeString"`
					} `json:"typeDescriptions"`
					TypeName struct {
						ID               int64  `json:"id"`
						Name             string `json:"name"`
						NodeType         string `json:"nodeType"`
						Src              string `json:"src"`
						TypeDescriptions struct {
							TypeIdentifier string `json:"typeIdentifier"`
							TypeString     string `json:"typeString"`
						} `json:"typeDescriptions"`
					} `json:"typeName"`
					Visibility string `json:"visibility"`
				} `json:"parameters"`
				Src string `json:"src"`
			} `json:"returnParameters"`
			Scope            int64  `json:"scope"`
			Src              string `json:"src"`
			StateMutability  string `json:"stateMutability"`
			Virtual          bool   `json:"virtual"`
			Visibility       string `json:"visibility"`
			TypeDescriptions struct {
				TypeIdentifier string `json:"typeIdentifier"`
				TypeString     string `json:"typeString"`
			} `json:"typeDescriptions"`
			TypeName struct {
				ID      int64 `json:"id"`
				KeyType struct {
					ID               int64  `json:"id"`
					Name             string `json:"name"`
					NodeType         string `json:"nodeType"`
					Src              string `json:"src"`
					TypeDescriptions struct {
						TypeIdentifier string `json:"typeIdentifier"`
						TypeString     string `json:"typeString"`
					} `json:"typeDescriptions"`
				} `json:"keyType"`
				Name             string `json:"name"`
				NodeType         string `json:"nodeType"`
				Src              string `json:"src"`
				StateMutability  string `json:"stateMutability"`
				TypeDescriptions struct {
					TypeIdentifier string `json:"typeIdentifier"`
					TypeString     string `json:"typeString"`
				} `json:"typeDescriptions"`
				ValueType struct {
					ID               int64  `json:"id"`
					Name             string `json:"name"`
					NodeType         string `json:"nodeType"`
					Src              string `json:"src"`
					TypeDescriptions struct {
						TypeIdentifier string `json:"typeIdentifier"`
						TypeString     string `json:"typeString"`
					} `json:"typeDescriptions"`
				} `json:"valueType"`
			} `json:"typeName"`
		} `json:"nodes"`
		Scope int64  `json:"scope"`
		Src   string `json:"src"`
	} `json:"nodes"`
	Src string `json:"src"`
}
