package pdk

type Plugin struct {
	// A unique id for this plugin. It may often contain the base id that extension points defined within the plugin
	// contain. For example mycompany.plugins.MyPlugin  and an extension point of this plugin might have an id of
	// mycompany.plugins.MyPlugin.EPName
	Id string `json:"id" yaml:"id"`

	// A more readable name of the plugin
	Name string `json:"name" yaml:"name"`

	// The version of this plugin
	Version string `json:"version" yaml:"version"`

	// This is the minimum version of the plugin engine that this plugin can work with.
	MinVersion string `json:"minVersion" yaml:"minVersion"`

	// A description of this plugin, what it does, provides, contributes to, etc.
	Description string `json:"description" yaml:"description"`

	// A slice of extension points that this plugin defines.. anchor points for other plugins to extend from
	ExtensionPoints []ExtensionPoint `json:"extensionPoints" yaml:"extensionPoints"`

	// A slice of extensions that attach to other plugin extension points.. contributions this plugin is adding to those
	// extension points.
	Extensions []Extension `json:"extensions" yaml:"extensions"`

	LoadOnStart bool `json:"loadOnStart" yaml:"loadOnStart"`
}
