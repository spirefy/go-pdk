package hostfuncs

import (
	"github.com/extism/go-pdk"
)

//go:wasmimport extism:host/pluginengine CallExtension
func callExtension(extensionId, data uint64) uint64

// CallExtension
//
// This function can be called by plugins to call one (or more) extensions matching the extension point id and version.
// If the slice of extensionsIds is nil or empty, the first extension found is called if the extension point match is
// found.
//
// This is a wrapper function which uses the imported callExtensionPointExtensions function implemented in the
// pluginengine plugin. This wrapper makes it easier for Go plugin developers to avoid the WASM memory management
func CallExtension(extensionId string, data []byte) ([]byte, error) {
	pdk.Log(pdk.LogDebug, "CallExtension")

	dta1 := pdk.AllocateString(extensionId)
	dta2 := pdk.AllocateBytes(data)

	off1 := dta1.Offset()
	off2 := dta2.Offset()

	resp := callExtension(off1, off2)

	mem1 := pdk.FindMemory(resp)
	bytes := mem1.ReadBytes()

	return bytes, nil
}
