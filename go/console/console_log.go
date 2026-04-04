package console

import (
	"encoding/json"
	"fmt"
)

// ConsoleLog imita o comportamento do JS, mostrando chaves e valores.
func ConsoleLog(v interface{}) {
	// O MarshalIndent cria a string JSON com recuo (espaços)
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		fmt.Printf("❌ [ConsoleLog Error]: %v\n", err)
		return
	}
	fmt.Println(string(b))
}
