// Code generated by 'option-gen'. DO NOT EDIT.

package kube

import (
	prompt "github.com/c-bata/go-prompt"
)

var attachOptions = []prompt.Suggest{
	prompt.Suggest{Text: "--container", Alias: "-c", Description: "Container name. If omitted, the first container in the pod will be chosen"},
	prompt.Suggest{Text: "--pod-running-timeout", Description: "The length of time (like 5s, 2m, or 3h, higher than zero) to wait until at least one pod is running"},
	prompt.Suggest{Text: "--stdin", Alias: "-i", Description: "Pass stdin to the container"},
	prompt.Suggest{Text: "--tty", Alias: "-t", Description: "Stdin is a TTY"},
}
