package main

import (
	def "github.com/ragulmathawa/go-storage/definitions"
	"github.com/ragulmathawa/go-storage/types"
)

var Metadata = def.Metadata{
	Name:  "us3",
	Infos: []def.Info{},
	Pairs: []def.Pair{},
	Factory: []def.Pair{
		def.PairWorkDir,
	},
	Service: def.Service{
		Features: types.ServiceFeatures{},
	},
	Storage: def.Storage{
		Features: types.StorageFeatures{},
	},
}
