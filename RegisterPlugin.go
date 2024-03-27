package hostfuncs

import (
	"encoding/json"
	"github.com/extism/go-pdk"
	"github.com/spirefy/go-pdk/types"
)

//go:wasmimport extism:host/user registerPlugin
func registerPlugin(d uint64) uint64

func RegisterPlugin(id, name, version, minVersion, description string, extensionPoints []types.ExtensionPoint, extensions []types.Extension) {
	// Call the host function

	plugin := types.Plugin{
		Id:              id,
		Name:            name,
		Version:         version,
		MinVersion:      minVersion,
		Description:     description,
		ExtensionPoints: extensionPoints,
		Extensions:      extensions,
	}

	dta, err := json.Marshal(plugin)
	if nil != err {
		pdk.Log(pdk.LogDebug, "Problem marshalling plugin registration struct: "+err.Error())
	}

	dataOffset := pdk.AllocateBytes(dta)

	offs1 := registerPlugin(dataOffset.Offset())

	mem1 := pdk.FindMemory(offs1)
	pdk.OutputMemory(mem1)
}
