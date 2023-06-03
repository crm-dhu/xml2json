package parser

import (
	"encoding/json"
	"fmt"
	"testing"
)

const xmlStr = `<?xml version="1.0" encoding="UTF-8"?>
<Flow xmlns="http://soap.sforce.com/2006/04/metadata">
    <actionCalls>
        <name>Send_Email</name>
        <label>Send Email</label>
        <actionName>emailSimple</actionName>
        <actionType>emailSimple</actionType>
        <flowTransactionModel>CurrentTransaction</flowTransactionModel>
        <inputParameters>
            <name>emailBody</name>
            <value>
                <stringValue>Congrats! Oppty is Closed Won</stringValue>
            </value>
        </inputParameters>
        <inputParameters>
            <name>emailAddresses</name>
            <value>
                <elementReference>$Record.Owner.Email</elementReference>
            </value>
        </inputParameters>
    </actionCalls>
    <processMetadataValues>
        <name>CanvasMode</name>
        <value>
            <stringValue>AUTO_LAYOUT_CANVAS</stringValue>
        </value>
    </processMetadataValues>
    <status>Draft</status>
    <status>Draft2</status>
</Flow>`

func TestFindTarget(t *testing.T) {
	d, _ := Xml2Map(xmlStr)
	// d := map[string]any{"a": 1, "b": map[string]any{"c": 2, "d": 3}, "e": map[string]any{"c": 4}}
	vals := find_target(d, "stringValue", []any{})
	fmt.Println(len(vals))
	for _, v := range vals {
		fmt.Println(v)
	}
}

func TestXml2Map(t *testing.T) {
	d, _ := Xml2Map(xmlStr)
	marshaled, _ := json.Marshal(d)
	fmt.Println(string(marshaled))
}
