package hostfuncs

import (
	"github.com/extism/go-pdk"
)

//go:wasmimport extism:host/user callExtension
func callExtension(d uint64) uint64

func CallExtension(data []byte) {
	// Call the host function

	pdk.Log(pdk.LogDebug, "Calling extension yah")
	dta := pdk.AllocateBytes(data)

	offs1 := callExtension(dta.Offset())

	mem1 := pdk.FindMemory(offs1)
	pdk.OutputMemory(mem1)
}
