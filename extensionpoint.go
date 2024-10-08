package pdk

type ExtensionPoint struct {
	// a unique identifier such as a name made up of the company, org, project and category separated by periods, or just
	// a simple string. The point is that this is to be matched by Extension's ExtensionPoint property in order for
	// extension to resolve to extension points when loaded.
	Id string `json:"id" yaml:"id"`

	// a meaningful description of this extension point that could be displayed in a plugin store for example. This
	// should probably provide details as to how the extension point will be called, when and what expectations if
	// any should be performed or provided by extensions
	Description string `json:"description" yaml:"description"`

	// a display or friendly name for this extension point, not to be confused with the Id.
	Name string `json:"name" yaml:"name"`

	// This version specified the minimum version the plugin engine must be for this extension point to be made
	// available
	Version string `json:"version" yaml:"version"`

	// This is a schema an extension point would potentially send to an extension when the plugin code makes a call to
	// one (or more) extensions.
	CallSchema map[string]any `json:"callSchema" yaml:"callSchema"`

	// This is a schema expected in response from an extension back to the extension point plugin code. Extension
	// implementations would need to return some subset of this structure in response to a call from an extension point.
	ResponseSchema map[string]any `json:"responseSchema" yaml:"responseSchema"`

	// This is an array of plugin Ids that must be available before this extension point can be used.
	Dependencies []Dependency `json:"dependencies" yaml:"dependencies"`
}

func CreateExtensionPoint(id, name, version, description string, dependencies []Dependency) *ExtensionPoint {
	if len(id) > 0 && len(name) > 0 {
		return &ExtensionPoint{
			Id:           id,
			Description:  description,
			Name:         name,
			Version:      version,
			Dependencies: dependencies,
		}
	}

	return nil
}
