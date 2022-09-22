// Code generated by jmx_codegen; DO NOT EDIT.

package jmaxml

import "time"

type VolcBody struct {
	Notice             *string                 `xml:"Notice" json:"notice,omitempty"`
	VolcanoInfos       []*VolcVolcanoInfo      `xml:"VolcanoInfo" json:"volcanoInfos,omitempty"`
	AshInfos           *VolcAshInfos           `xml:"AshInfos" json:"ashInfos,omitempty"`
	VolcanoInfoContent *VolcVolcanoInfoContent `xml:"VolcanoInfoContent" json:"volcanoInfoContent,omitempty"`
	VolcanoObservation *VolcVolcanoObservation `xml:"VolcanoObservation" json:"volcanoObservation,omitempty"`
	Text               *string                 `xml:"Text" json:"text,omitempty"`
}

type VolcVolcanoInfo struct {
	Type  string      `xml:"type,attr" json:"type"`
	Items []*VolcItem `xml:"Item" json:"items"`
}

type VolcItem struct {
	EventTime *VolcEventTime `xml:"EventTime" json:"eventTime,omitempty"`
	Kind      VolcKind       `xml:"Kind" json:"kind"`
	LastKind  *VolcKind      `xml:"LastKind" json:"lastKind,omitempty"`
	Areas     VolcAreas      `xml:"Areas" json:"areas"`
}

type VolcEventTime struct {
	EventDateTime        *EbDateTime `xml:"EventDateTime" json:"eventDateTime,omitempty"`
	EventDateTimeUTC     *EbDateTime `xml:"EventDateTimeUTC" json:"eventDateTimeUTC,omitempty"`
	EventDateTimeComment *string     `xml:"EventDateTimeComment" json:"eventDateTimeComment,omitempty"`
}

type VolcKind struct {
	Name       string        `xml:"Name" json:"name"`
	FormalName *string       `xml:"FormalName" json:"formalName,omitempty"`
	Code       *string       `xml:"Code" json:"code,omitempty"`
	Condition  *string       `xml:"Condition" json:"condition,omitempty"`
	Property   *VolcProperty `xml:"Property" json:"property,omitempty"`
}

type VolcAreas struct {
	CodeType string      `xml:"codeType,attr" json:"codeType"`
	Areas    []*VolcArea `xml:"Area" json:"areas"`
}

type VolcArea struct {
	Name             string        `xml:"Name" json:"name"`
	Code             string        `xml:"Code" json:"code"`
	Coordinate       *EbCoordinate `xml:"Coordinate" json:"coordinate,omitempty"`
	AreaFromMark     *string       `xml:"AreaFromMark" json:"areaFromMark,omitempty"`
	CraterName       *string       `xml:"CraterName" json:"craterName,omitempty"`
	CraterCoordinate *EbCoordinate `xml:"CraterCoordinate" json:"craterCoordinate,omitempty"`
}

type VolcAshInfos struct {
	Type     string         `xml:"type,attr" json:"type"`
	AshInfos []*VolcAshInfo `xml:"AshInfo" json:"ashInfos"`
}

type VolcAshInfo struct {
	Type      string      `xml:"type,attr" json:"type"`
	StartTime time.Time   `xml:"StartTime" json:"startTime"`
	EndTime   time.Time   `xml:"EndTime" json:"endTime"`
	Items     []*VolcItem `xml:"Item" json:"items"`
}

type VolcProperty struct {
	Size           *VolcSize        `xml:"Size" json:"size,omitempty"`
	Polygons       []*EbCoordinate  `xml:"Polygon" json:"polygons"`
	PlumeDirection EbPlumeDirection `xml:"PlumeDirection" json:"plumeDirection"`
	Distance       VolcDistance     `xml:"Distance" json:"distance"`
	Remark         *string          `xml:"Remark" json:"remark,omitempty"`
}

type VolcSize struct {
	Value NullableFloat `xml:",chardata"`
	Type  *string       `xml:"type,attr" json:"type,omitempty"`
	Unit  *string       `xml:"unit,attr" json:"unit,omitempty"`
}

type VolcDistance struct {
	Value       string  `xml:",chardata"`
	Type        *string `xml:"type,attr" json:"type,omitempty"`
	Unit        *string `xml:"unit,attr" json:"unit,omitempty"`
	Description *string `xml:"description,attr" json:"description,omitempty"`
}

type VolcVolcanoInfoContent struct {
	VolcanoHeadline   *string `xml:"VolcanoHeadline" json:"volcanoHeadline,omitempty"`
	VolcanoActivity   *string `xml:"VolcanoActivity" json:"volcanoActivity,omitempty"`
	VolcanoPrevention *string `xml:"VolcanoPrevention" json:"volcanoPrevention,omitempty"`
	NextAdvisory      *string `xml:"NextAdvisory" json:"nextAdvisory,omitempty"`
	OtherInfo         *string `xml:"OtherInfo" json:"otherInfo,omitempty"`
	Appendix          *string `xml:"Appendix" json:"appendix,omitempty"`
	Text              *string `xml:"Text" json:"text,omitempty"`
}

type VolcVolcanoObservation struct {
	EventTime        *VolcEventTime `xml:"EventTime" json:"eventTime,omitempty"`
	ColorPlume       *VolcPlume     `xml:"ColorPlume" json:"colorPlume,omitempty"`
	WhitePlume       *VolcPlume     `xml:"WhitePlume" json:"whitePlume,omitempty"`
	OtherObservation *string        `xml:"OtherObservation" json:"otherObservation,omitempty"`
	Appendix         *string        `xml:"Appendix" json:"appendix,omitempty"`
}

type VolcPlume struct {
	PlumeHeightAboveCrater   EbPlumeHeight    `xml:"PlumeHeightAboveCrater" json:"plumeHeightAboveCrater"`
	PlumeHeightAboveSeaLevel *EbPlumeHeight   `xml:"PlumeHeightAboveSeaLevel" json:"plumeHeightAboveSeaLevel,omitempty"`
	PlumeDirection           EbPlumeDirection `xml:"PlumeDirection" json:"plumeDirection"`
	PlumeComment             *string          `xml:"PlumeComment" json:"plumeComment,omitempty"`
}
