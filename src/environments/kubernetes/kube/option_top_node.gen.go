// Code generated by 'option-gen'. DO NOT EDIT.

package kube

import (
	prompt "github.com/c-bata/go-prompt"
)

var topNodeOptions = []prompt.Suggest{
	prompt.Suggest{Text: "--heapster-namespace", Description: "Namespace Heapster service is located in"},
	prompt.Suggest{Text: "--heapster-port", Description: "Port name in service to use"},
	prompt.Suggest{Text: "--heapster-scheme", Description: "Scheme (http or https) to connect to Heapster as"},
	prompt.Suggest{Text: "--heapster-service", Description: "Name of Heapster service"},
	prompt.Suggest{Text: "--selector", Alias: "-l", Description: "Selector (label query) to filter on, supports '=', '==', and '!='.(e.g. -l key1=value1,key2=value2)"},
}
