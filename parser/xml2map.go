package parser

import (
	"encoding/json"
	"fmt"

	"github.com/beevik/etree"
)

type strMap map[string]any

func ParseXML(xmlStr string) (err error) {

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

	var d strMap
	doc := etree.NewDocument()
	doc.ReadFromString(xmlStr)
	root := doc.SelectElement("Flow")

	if err = recursiveParse(root, d, array_keys); err != nil {
		fmt.Println("Something is wrong")
		return nil
	}
	y, _ := json.Marshal(d)
	fmt.Println(string(y))
	return nil

}

// func ParseXML(xmlStr string) {
// 	x := map[string]any{"a": 1, "b": map[string]any{"c": 2}, "d": []int{1, 2, 3}}
// 	y, _ := json.Marshal(x)
// 	fmt.Println(string(y))
// 	doc := etree.NewDocument()
// 	doc.ReadFromString(xmlStr)
// 	root := doc.SelectElement("Flow")
// 	for _, e := range root.ChildElements() {
// 		fmt.Println(e.Tag)
// 		if len(e.ChildElements()) == 0 {
// 			fmt.Println(e.Text())
// 		}
// 	}
// }

func recursiveParse(root *etree.Element, d strMap, array_keys *Set) (err error) {
	if root == nil {
		return nil
	}
	children := root.ChildElements()
	if len(children) == 1 {
		if array_keys.Contains(root.Tag) {
			if _, ok := d[root.Tag]; ok {
				d[root.Tag] = append(d[root.Tag].([]string), root.Text())
			} else {
				d[root.Tag] = []string{root.Text()}

			}
		} else {
			d[root.Tag] = root.Text()
		}
		return nil
	}
	var d1 strMap
	for _, child := range children {
		recursiveParse(child, d1, array_keys)
	}
	if array_keys.Contains(root.Tag) {
		if _, ok := d[root.Tag]; ok {
			d[root.Tag] = append(d[root.Tag].([]any), d1)
		} else {
			d[root.Tag] = []any{d1}

		}
	} else {
		d[root.Tag] = d1
	}
	return nil

}
