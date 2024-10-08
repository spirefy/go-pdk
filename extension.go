package pdk

type Extension struct {
	// a unique identifier such as a name made up of the company, org, project and category separated by periods, or just
	// a simple string.
	Id string `json:"id" yaml:"id"`

	// the extension point Id that this extension provides functionality to/for.
	ExtensionPoint string `json:"extensionPoint" yaml:"extensionPoint"`

	// a meaningful description of this extension that could be displayed in a plugin store for example
	Description string `json:"description" yaml:"description"`

	// a display or friendly name for this extension, not to be confused with the Id.
	Name string `json:"name" yaml:"name"`

	// a func IN the WASM plugin that matches this value that would be called by the extension point using the plugin
	// engine's host function to call extension functions.
	Func string `json:"func" yaml:"func"`

	// This property can hold custom data for an extension point to use without having to call the extension func
	// to have data returned. This is useful when an ExtensionPoint want's to build up say a Menu system or a Help
	// system that is composed of static data and does not require the Extension func to execute to return such data.
	// Each ExtensionPoint would define the format (structure) that this metadata would need to be in and it is up
	// to the exported register() function to marshal the data into a []byte before returning the Extension(s) as
	// part of the register return []byte data. ExtensionPoint funcs would be called to process the metadata however
	// necessary.
	MetaData map[string]any `json:"metadata" yaml:"metadata"`

	// This is a slice of dependencies this extension depends on. If provided, the extension, extension point or
	// plugin in each dependency MUST resolve before the extension is usable (resolved)
	Dependencies []Dependency `json:"dependencies" yaml:"dependencies"`
}

func CreateExtension(id, name, extensionPoint, description, funcName string, metadata map[string]any, dependencies []Dependency) *Extension {
	if len(id) > 0 && len(name) > 0 && len(extensionPoint) > 0 && len(funcName) > 0 {
		return &Extension{
			Id:             id,
			ExtensionPoint: extensionPoint,
			Description:    description,
			Name:           name,
			Func:           funcName,
			MetaData:       metadata,
			Dependencies:   dependencies,
		}
	}

	return nil
}
