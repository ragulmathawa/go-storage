package main

import (
	def "github.com/ragulmathawa/go-storage/definitions"
)

func main() {
	def.GenerateService(Metadata, "generated.go")
}
