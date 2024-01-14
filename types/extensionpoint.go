package types

type ExtensionPoint struct {
	// a unique identifier such as a name made up of the company, org, project and category separated by periods, or just
	// a simple string. The point is that this is to be matched by Extension's ExtensionPoint property in order for
	// extension to resolve to extension points when loaded.
	Id string `json:"id"`

	// a meaningful description of this extension point that could be displayed in a plugin store for example. This
	// should probably provide details as to how the extension point will be called, when and what expectations if
	// any should be performed or provided by extensions
	Description string `json:"description"`

	// a display or friendly name for this extension point, not to be confused with the Id.
	Name string `json:"name"`

	// This version specified the minimum version the plugin engine must be for this extension point to be made
	// available
	Version string `json:"version"`

	// This is an array of plugin Ids that must be available before this extension point can be used.
	Dependencies []string `json:"dependencies"`

	// If this extension point represents an event trigger this is the event that will trigger it. Any extensions
	// matched to this extension point will be expected to support the event via an event handler function
	Event string `json:"event"`

	// This is either a json schema as a string or a chunk of json representing the full payload structure that this
	// extension point would send to any resolved extensions. This is mostly for display purposes or informational
	// purposes as extension point documentation should include info on the particular payload dynamics if any are
	// present with details on the various properties, etc. This particular property would be useful in say, a plugin
	// store that shows the structure of the extension point and displays this schema for informational purpose.
	Schema string `json:"schema"`
}
