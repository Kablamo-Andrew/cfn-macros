package transform

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

var sampleJson = `{
  "Parameters":{},
  "Mappings":{},
  "Resources": {
    "MyVolumeAttachment": {
      "Type": "AWS::EC2::VolumeAttachment",
      "Properties": {
        "Device": "/dev/sda",
        "InstanceId": {
          "Name": "MyInstance",
		  "Kind": "Resource",
          "Type": "AWS::EC2::Instance",
          "Properties": {}
        },
        "VolumeId": {
          "Name": "MyVolume",
		  "Kind": "Resource",
          "Type": "AWS::EC2::Volume",
          "Properties": {
            "Size": {
				"Name": "MyVolumeSize",
				"Kind": "Parameter",
				"Type": "Number",
				"Default": 100
			}
          }
        }
      }
    }
  },
  "Outputs":{}
}`

func TestSimpleTransform(t *testing.T) {
	var templateObj map[string]interface{}

	err := json.Unmarshal([]byte(sampleJson), &templateObj)
	assert.Nil(t, err)

	transformed := Transform(templateObj)
	assert.Len(t, transformed["Parameters"].(map[string]interface{}), 1)
	assert.Len(t, transformed["Resources"].(map[string]interface{}), 3)
}

var sampleWithArray = `{
  "Parameters":{},
  "Mappings":{},
  "Resources": {
    "MyVolumeAttachment": {
      "Type": "AWS::EC2::VolumeAttachment",
      "Properties": {
        "Device": "/dev/sda",
        "InstanceId": {
          "Name": "MyInstance",
		  "Kind": "Resource",
          "Type": "AWS::EC2::Instance",
          "Properties": {}
        },
        "VolumeIds": [
		  {
            "Name": "MyVolume1",
		    "Kind": "Resource",
            "Type": "AWS::EC2::Volume",
            "Properties": {
              "Size": 100
            }
          },
		  {
            "Name": "MyVolume2",
		    "Kind": "Resource",
            "Type": "AWS::EC2::Volume",
            "Properties": {
              "Size": 100
            }
          }
        ]
      }
    }
  },
  "Outputs":{}
}`

func TestArrayTransform(t *testing.T) {
	var templateObj map[string]interface{}

	err := json.Unmarshal([]byte(sampleWithArray), &templateObj)
	assert.Nil(t, err)

	transformed := Transform(templateObj)
	assert.Len(t, transformed["Resources"].(map[string]interface{}), 4)
}
