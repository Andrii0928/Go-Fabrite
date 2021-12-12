package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRelativePathToGitComponent(t *testing.T) {
	subcomponent := Component{
		Name:   "efk",
		Method: "git",
		Source: "https://github.com/microsoft/fabrikate-elasticsearch-fluentd-kibana",
	}

	assert.Equal(t, subcomponent.RelativePathTo(), "components/efk")
}

func TestRelativePathToDirectoryComponent(t *testing.T) {
	subcomponent := Component{
		Name:   "infra",
		Source: "./infra",
	}

	assert.Equal(t, subcomponent.RelativePathTo(), "./infra")
}

func TestLoadComponent(t *testing.T) {
	component := Component{
		PhysicalPath: "../../testdata/definition/infra",
		LogicalPath:  "infra",
	}

	component, err := component.LoadComponent()
	assert.Nil(t, err)

	assert.Nil(t, err)
	assert.Equal(t, component.Name, "infra")
	assert.Equal(t, len(component.Subcomponents), 1)
	assert.Equal(t, component.Subcomponents[0].Name, "efk")
	assert.Equal(t, component.Subcomponents[0].Source, "https://github.com/microsoft/fabrikate-elasticsearch-fluentd-kibana")
	assert.Equal(t, component.Subcomponents[0].Method, "git")
}

func TestLoadBadYAMLComponent(t *testing.T) {
	component := Component{
		PhysicalPath: "../../testdata/badyamldefinition",
		LogicalPath:  "",
	}

	component, err := component.LoadComponent()
	assert.NotNil(t, err)
}

func TestLoadBadJSONComponent(t *testing.T) {
	component := Component{
		PhysicalPath: "../../testdata/badjsondefinition",
		LogicalPath:  "",
	}

	component, err := component.LoadComponent()
	assert.NotNil(t, err)
}

func TestLoadConfig(t *testing.T) {
	component := Component{
		PhysicalPath: "../../testdata/generate/infra",
		LogicalPath:  "infra",
	}

	component, err := component.LoadComponent()
	assert.Nil(t, err)

	err = component.LoadConfig([]string{"prod-east", "prod"})

	assert.Nil(t, err)
}

func TestUpdateRootComponentPath(t *testing.T) {
	component := Component{
		PhysicalPath: "../../testdata/definition/infra-single",
		LogicalPath:  "infra-single",
	}

	component, err := component.LoadComponent()
	assert.Nil(t, err)

	err = component.LoadConfig([]string{})
	assert.Nil(t, err)

	assert.Equal(t, 0, len(component.Subcomponents))

	component, err = component.UpdateComponentPath("../../testdata/definition/infra-single", []string{})
	assert.Nil(t, err)
	assert.Equal(t, 4, len(component.Subcomponents))
}

func TestIteratingDefinition(t *testing.T) {
	callbackCount := 0

	rootInit := func(startPath string, environments []string, c Component) (component Component, err error) {
		return c, nil
	}

	results := WalkComponentTree("../../testdata/iterator", []string{""}, func(path string, component *Component) (err error) {
		callbackCount++
		return nil
	}, rootInit)

	var err error
	components := make([]Component, 0)
	for result := range results {
		if result.Error != nil {
			err = result.Error
		} else if result.Component != nil {
			components = append(components, *result.Component)
		}
	}

	assert.Nil(t, err)
	assert.Equal(t, 3, len(components))
	assert.Equal(t, callbackCount, len(components))

	assert.Equal(t, components[1].PhysicalPath, "../../testdata/iterator/infra")
	assert.Equal(t, components[1].LogicalPath, "infra")

	assert.Equal(t, components[2].PhysicalPath, "../../testdata/iterator/infra/components/efk")
	assert.Equal(t, components[2].LogicalPath, "infra/efk")
}

func TestWriteComponent(t *testing.T) {
	component := Component{
		PhysicalPath: "../../testdata/install",
		LogicalPath:  "",
	}

	component, err := component.LoadComponent()
	assert.Nil(t, err)

	err = component.Write()
	assert.Nil(t, err)
}

func TestLoadDisabledComponentDefaultValue(t *testing.T) {
	component := Component{
		PhysicalPath: "../../testdata/disabled",
		LogicalPath:  "",
	}

	component, err := component.LoadComponent()
	assert.Nil(t, err)

	err = component.LoadConfig([]string{"default"})
	assert.Nil(t, err)

	assert.Equal(t, false, component.Config.Subcomponents["cloud-native"].Disabled)
	assert.Equal(t, false, component.Config.Subcomponents["elasticsearch"].Disabled)

}

func TestLoadDisabledComponent(t *testing.T) {
	component := Component{
		PhysicalPath: "../../testdata/disabled",
		LogicalPath:  "",
	}

	component, err := component.LoadComponent()
	assert.Nil(t, err)

	err = component.LoadConfig([]string{"disabled"})
	assert.Nil(t, err)

	assert.Equal(t, true, component.Config.Subcomponents["cloud-native"].Disabled)
	assert.Equal(t, true, component.Config.Subcomponents["elasticsearch"].Disabled)
}
