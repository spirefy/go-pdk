package hostfuncs

import (
	"github.com/extism/go-pdk"
	"github.com/spirefy/go-pdk/types"
	"strings"
)

//go:wasmimport extism:host/user getExtensionsForExtensionPoint
func getExtensionsForExtensionPoint(ep, ver uint64) uint64

func GetExtensionsForExtensionPoint(ep string, versions []string) ([]types.Extension, error) {
	// Call the host function

	pdk.Log(pdk.LogDebug, "Calling getExtensionsForExtensionPoint")
	dta1 := pdk.AllocateString(ep)

	joinedVersions := strings.Join(versions, ",")

	dta2 := pdk.AllocateBytes([]byte(joinedVersions))

	off := dta1.Offset()
	off2 := dta2.Offset()
	offs1 := getExtensionsForExtensionPoint(off, off2)

	mem1 := pdk.FindMemory(offs1)

	pdk.OutputMemory(mem1)

	return nil, nil
}
