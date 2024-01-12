package hostfuncs

import (
	"github.com/extism/go-pdk"
)

//go:wasmimport extism:host/user sendEvent
func sendEvent(e, d uint64) uint64

func SendEvent(event string, data []byte) {
	// Call the host function

	evt := pdk.AllocateString(event)
	dta := pdk.AllocateBytes(data)

	offs1 := sendEvent(evt.Offset(), dta.Offset())

	mem1 := pdk.FindMemory(offs1)
	pdk.OutputMemory(mem1)
}
