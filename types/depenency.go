package types

type Dependency struct {
	*ExtensionPoint
	*Extension
	*Plugin
}

func CreateDependency(ep *ExtensionPoint, e *Extension, p *Plugin) *Dependency {
	if nil != ep || nil != e || nil != p {
		return &Dependency{
			ExtensionPoint: ep,
			Extension:      e,
			Plugin:         p,
		}
	}

	return nil
}
