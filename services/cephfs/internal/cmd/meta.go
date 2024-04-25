package main

import (
	def "github.com/ragulmathawa/go-storage/definitions"
	"github.com/ragulmathawa/go-storage/types"
)

var Metadata = def.Metadata{
	Name:  "cephfs",
	Pairs: []def.Pair{},
	Infos: []def.Info{},
	Factory: []def.Pair{
		def.PairWorkDir,
	},
	Service: def.Service{},
	Storage: def.Storage{
		Features: types.StorageFeatures{
			WriteEmptyObject: true,

			Create:   true,
			Delete:   true,
			List:     true,
			Metadata: true,
			Read:     true,
			Write:    true,
		},
	},
}
