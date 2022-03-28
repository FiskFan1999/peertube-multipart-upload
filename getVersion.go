package main

import (
	_ "embed"
	"encoding/json"
	"log"
)

//go:embed version
var VersionFile []byte

type VersStrType struct {
	Version string
	Commit  string
}

var Version *VersStrType

func GetVersStr(f []byte) (out *VersStrType, err error) {
	out = new(VersStrType)
	err = json.Unmarshal(f, out)
	return
}

func init() {
	var err error
	Version, err = GetVersStr(VersionFile)
	if err != nil {
		log.Fatalf("Error when reading from Version file: \"%+v\"\n", err)
	}
}
