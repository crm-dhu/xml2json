package parser

import (
	"errors"

	"github.com/beevik/etree"
)

func find_target(jsonElement any, lookup_key string, values []any) []any {
	switch v := jsonElement.(type) {
	case map[string]any:
		for key, value := range v {
			if key == lookup_key {
				values = append(values, value)
			} else {
				values = find_target(value, lookup_key, values)
			}
		}

	case []any:
		for _, e := range v {
			values = find_target(e, lookup_key, values)
		}
	}
	return values

}

func Xml2Map(xmlStr string) (jsonMap map[string]any, err error) {

	// data_elements := NewSet("recordLookups", "recordUpdates", "recordCreates", "recordDeletes")
	// interaction_elements := NewSet("actionCalls", "subflows", "screens")
	// logic_elements := NewSet("decisions", "assignments", "waits", "loops", "collectionProcessors")
	// elements := NewSet()
	// elements.AddSet(data_elements)
	// elements.AddSet(interaction_elements)
	// elements.AddSet(logic_elements)
	// resources := NewSet("variables", "constants", "formulas", "textTemplates", "stages")
	array_keys := NewSet("actionCalls", "apexPluginCalls", "assignments", "choices", "collectionProcessors",
		"constants", "decisions", "dynamicChoiceSets", "formulas", "loops", "orchestratedStages",
		"processMetadataValues", "recordCreates", "recordDeletes", "recordLookups",
		"recordRollbacks", "recordUpdates", "screens", "stages", "steps", "subflows",
		"textTemplates", "variables", "waits", "dataTypeMappings", "inputParameters",
		"outputParameters", "assignmentItems", "conditions", "mapItems", "sortOptions",
		"rules", "filters", "outputAssignments", "processMetadataValues", "exitActionInputParameters",
		"exitActionOutputParameters", "exitConditions", "stageSteps", "inputAssignments",
		"choiceReferences", "assignees", "entryActionInputParameters", "entryActionOutputParameters",
		"entryConditions", "scheduledPaths", "connectors", "waitEvents")

	d := map[string]any{}
	doc := etree.NewDocument()
	doc.ReadFromString(xmlStr)
	root := doc.SelectElement("Flow")

	if err = recursiveParse(root, d, array_keys); err != nil {
		return nil, errors.New("parsing failed")
	}
	return d, nil
}

func recursiveParse(root *etree.Element, d map[string]any, array_keys *Set) (err error) {
	if root == nil {
		return nil
	}
	children := root.ChildElements()
	if len(children) == 0 {
		if array_keys.Contains(root.Tag) {
			if _, ok := d[root.Tag]; ok {
				if leaves, ok := d[root.Tag].([]string); ok {
					d[root.Tag] = append(leaves, root.Text())
				} else {
					errors.New("leaves cannot be converted.")
				}
			} else {
				d[root.Tag] = []string{root.Text()}
			}
		} else {
			d[root.Tag] = root.Text()
		}
		return nil
	}
	d1 := map[string]any{}
	for _, child := range children {
		recursiveParse(child, d1, array_keys)
	}
	if array_keys.Contains(root.Tag) {
		if _, ok := d[root.Tag]; ok {
			if nodes, ok := d[root.Tag].([]any); ok {
				d[root.Tag] = append(nodes, d1)
			} else {
				errors.New("nodes cannot be converted.")
			}
		} else {
			d[root.Tag] = []any{d1}
		}
	} else {
		d[root.Tag] = d1
	}
	return nil

}
