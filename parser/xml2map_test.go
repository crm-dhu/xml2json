package parser

import (
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

func TestParseXML(t *testing.T) {
	ParseXML(xmlStr)
}
