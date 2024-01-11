package hostfuncs

//go:wasmimport extism:host/user sendevent
func sendevent(e, d uint64) uint64

func SendEvent(event string, data []byte) {
	// Call the host function
	/*
		evt := pdk.AllocateString(event)
		dta := pdk.AllocateBytes(data)


		offs1 := sendevent(evt.Offset(), dta.Offset())


		mem1 := pdk.FindMemory(offs1)

		buffMem1 := make([]byte, mem1.Length())
		mem1.Load(buffMem1)

	*/
}
