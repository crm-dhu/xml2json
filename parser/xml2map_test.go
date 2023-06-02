package parser

import (
	"log"
	"testing"
)

const (
    xmlWithHeader = `<?xml version="1.0" encoding="UTF-8"?>
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
</Flow>`

xmlStrWithoutHeader = `<Flow xmlns="http://soap.sforce.com/2006/04/metadata">
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
</Flow>`
)

func TestGetXmlMap(t *testing.T) {
	receivingXmlMap, err := GetXmlMap(xmlWithHeader, "->")
	if err != nil {
		log.Println("Err: ", err)
		return
	} else if receivingXmlMap == nil {
		log.Println("Received nil xml map")
		return
	}
	for key, value := range receivingXmlMap {
		log.Println(key, " > ", value)
	}
}

func TestGetXmlMap2(t *testing.T) {
	receivingXmlMap, err := GetXmlMap(xmlStrWithoutHeader, "->")
	if err != nil {
		log.Println("Err: ", err)
		return
	} else if receivingXmlMap == nil {
		log.Println("Received nil xml map")
		return
	}
	for key, value := range receivingXmlMap {
		log.Println(key, " > ", value)
	}
}
