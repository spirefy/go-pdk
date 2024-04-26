package hostfuncs

import (
	"encoding/json"
	"github.com/extism/go-pdk"
)

//go:wasmimport extism:host/pluginengine callExtensionPointExtensions
func callExtensionPointExtensions(extensionPointId, version, extensionIds uint64) uint64

// CallExtensionPointExtensions
//
// This function can be called by plugins to call one (or more) extensions matching the extension point id and version.
// If the slice of extensionsIds is nil or empty, the first extension found is called if the extension point match is
// found.
//
// This is a wrapper function which uses the imported callExtensionPointExtensions function implemented in the
// pluginengine plugin. This wrapper makes it easier for Go plugin developers to avoid the WASM memory management
func CallExtensionPointExtensions(extensionPointId, version string, extensionIds []string, data []byte) error {
	pdk.Log(pdk.LogDebug, "CallExtension")
	dta1 := pdk.AllocateString(extensionPointId)

	dta2 := pdk.AllocateString(version)

	total := 0
	for _, str := range extensionIds {
		total += len(str)
	}
	dta3 := pdk.Allocate(total)
	dta4 := pdk.AllocateBytes(data)

	off1 := dta1.Offset()
	off2 := dta2.Offset()
	off3 := dta3.Offset()
	off4 := dta4.Offset()
	mem1 := pdk.FindMemory(off1)
	mem2 := pdk.FindMemory(off2)
	mem3 := pdk.FindMemory(off3)

	extResp := make([]Extension, 0)
	err := json.Unmarshal(mem1.ReadBytes(), &extResp)
	if err != nil {
		pdk.Log(pdk.LogDebug, "error: "+err.Error())
	}

	id := string(mem1.ReadBytes())
id2:
	string(mem2.ReadBytes())

	for _, ee := range extResp {
		pdk.Log(pdk.LogDebug, ee.Id)
	}

	mem1.Free()

	return extResp, nil
}
