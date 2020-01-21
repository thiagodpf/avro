// Code generated by avrogen. DO NOT EDIT.

package cloudEvent

import (
	"github.com/heetch/avro/avrotypegen"
)

type Metadata struct {
	Id     string `json:"id"`
	Source string `json:"source"`
	Time   int64  `json:"time"`
}

// AvroRecord implements the avro.AvroRecord interface.
func (Metadata) AvroRecord() avrotypegen.RecordInfo {
	return avrotypegen.RecordInfo{
		Schema: `{"fields":[{"name":"id","type":"string"},{"name":"source","type":"string"},{"name":"time","type":"long"}],"name":"Metadata","namespace":"avro.apache.org","type":"record"}`,
		Required: []bool{
			0: true,
			1: true,
			2: true,
		},
	}
}

// TODO implement MarshalBinary and UnmarshalBinary methods?

type SomeEvent struct {
	Metadata Metadata
}

// AvroRecord implements the avro.AvroRecord interface.
func (SomeEvent) AvroRecord() avrotypegen.RecordInfo {
	return avrotypegen.RecordInfo{
		Schema: `{"fields":[{"name":"Metadata","type":{"fields":[{"name":"id","type":"string"},{"name":"source","type":"string"},{"name":"time","type":"long"}],"name":"Metadata","namespace":"avro.apache.org","type":"record"}}],"name":"SomeEvent","namespace":"bar","type":"record"}`,
		Required: []bool{
			0: true,
		},
	}
}

// TODO implement MarshalBinary and UnmarshalBinary methods?
