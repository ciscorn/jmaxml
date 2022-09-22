// Code generated by jmx_codegen; DO NOT EDIT.

package jmaxml

import "time"

type EbDateTime struct {
	Value       time.Time `xml:",chardata"`
	Type        *string   `xml:"type,attr" json:"type,omitempty"`
	Significant *string   `xml:"significant,attr" json:"significant,omitempty"`
	Precision   *string   `xml:"precision,attr" json:"precision,omitempty"`
	Dubious     *string   `xml:"dubious,attr" json:"dubious,omitempty"`
	Description *string   `xml:"description,attr" json:"description,omitempty"`
}

type EbCoordinate struct {
	Value       string  `xml:",chardata"`
	Type        *string `xml:"type,attr" json:"type,omitempty"`
	Datum       *string `xml:"datum,attr" json:"datum,omitempty"`
	Condition   *string `xml:"condition,attr" json:"condition,omitempty"`
	Description *string `xml:"description,attr" json:"description,omitempty"`
}

type EbCircle struct {
	Type       *string         `xml:"type,attr" json:"type,omitempty"`
	BasePoints []*EbCoordinate `xml:"BasePoint" json:"basePoints,omitempty"`
	Axes       *EbAxes         `xml:"Axes" json:"axes,omitempty"`
}

type EbAxes struct {
	Axes      []*EbAxis `xml:"Axis" json:"axes,omitempty"`
	LongAxes  []*EbAxis `xml:"LongAxis" json:"longAxes,omitempty"`
	ShortAxes []*EbAxis `xml:"ShortAxis" json:"shortAxes,omitempty"`
}

type EbAxis struct {
	Directions []*EbDirection `xml:"Direction" json:"directions,omitempty"`
	Bearings   []*EbBearings  `xml:"Bearings" json:"bearingses,omitempty"`
	Radiuses   []*EbRadius    `xml:"Radius" json:"radiuses"`
}

type EbPressure struct {
	Value       NullableFloat `xml:",chardata"`
	Type        string        `xml:"type,attr" json:"type"`
	Unit        *string       `xml:"unit,attr" json:"unit,omitempty"`
	RefID       *uint8        `xml:"refID,attr" json:"refID,omitempty"`
	Condition   *string       `xml:"condition,attr" json:"condition,omitempty"`
	Description *string       `xml:"description,attr" json:"description,omitempty"`
}

type EbTemperature struct {
	Value       NullableFloat `xml:",chardata"`
	Type        string        `xml:"type,attr" json:"type"`
	Unit        *string       `xml:"unit,attr" json:"unit,omitempty"`
	RefID       *uint8        `xml:"refID,attr" json:"refID,omitempty"`
	Condition   *string       `xml:"condition,attr" json:"condition,omitempty"`
	Description *string       `xml:"description,attr" json:"description,omitempty"`
}

type EbHumidity struct {
	Value       NullableFloat `xml:",chardata"`
	Type        string        `xml:"type,attr" json:"type"`
	Unit        *string       `xml:"unit,attr" json:"unit,omitempty"`
	RefID       *uint8        `xml:"refID,attr" json:"refID,omitempty"`
	Condition   *string       `xml:"condition,attr" json:"condition,omitempty"`
	Description *string       `xml:"description,attr" json:"description,omitempty"`
}

type EbWindDirection struct {
	Value       string  `xml:",chardata"`
	Type        string  `xml:"type,attr" json:"type"`
	Unit        *string `xml:"unit,attr" json:"unit,omitempty"`
	RefID       *uint8  `xml:"refID,attr" json:"refID,omitempty"`
	Condition   *string `xml:"condition,attr" json:"condition,omitempty"`
	Description *string `xml:"description,attr" json:"description,omitempty"`
}

type EbWindDegree struct {
	Value       NullableFloat `xml:",chardata"`
	Type        string        `xml:"type,attr" json:"type"`
	Unit        *string       `xml:"unit,attr" json:"unit,omitempty"`
	RefID       *uint8        `xml:"refID,attr" json:"refID,omitempty"`
	Condition   *string       `xml:"condition,attr" json:"condition,omitempty"`
	Description *string       `xml:"description,attr" json:"description,omitempty"`
}

type EbWindSpeed struct {
	Value       NullableFloat `xml:",chardata"`
	Type        string        `xml:"type,attr" json:"type"`
	Unit        *string       `xml:"unit,attr" json:"unit,omitempty"`
	RefID       *uint8        `xml:"refID,attr" json:"refID,omitempty"`
	Condition   *string       `xml:"condition,attr" json:"condition,omitempty"`
	Description *string       `xml:"description,attr" json:"description,omitempty"`
}

type EbWindScale struct {
	Value       NullableInteger `xml:",chardata"`
	Type        string          `xml:"type,attr" json:"type"`
	Unit        *string         `xml:"unit,attr" json:"unit,omitempty"`
	RefID       *uint8          `xml:"refID,attr" json:"refID,omitempty"`
	Condition   *string         `xml:"condition,attr" json:"condition,omitempty"`
	Description *string         `xml:"description,attr" json:"description,omitempty"`
}

type EbSunshine struct {
	Value       NullableFloat `xml:",chardata"`
	Type        string        `xml:"type,attr" json:"type"`
	Unit        *string       `xml:"unit,attr" json:"unit,omitempty"`
	RefID       *uint8        `xml:"refID,attr" json:"refID,omitempty"`
	Condition   *string       `xml:"condition,attr" json:"condition,omitempty"`
	Description *string       `xml:"description,attr" json:"description,omitempty"`
}

type EbPrecipitation struct {
	Value       NullableFloat `xml:",chardata"`
	Type        string        `xml:"type,attr" json:"type"`
	Unit        *string       `xml:"unit,attr" json:"unit,omitempty"`
	RefID       *uint8        `xml:"refID,attr" json:"refID,omitempty"`
	Condition   *string       `xml:"condition,attr" json:"condition,omitempty"`
	Description *string       `xml:"description,attr" json:"description,omitempty"`
}

type EbProbabilityOfPrecipitation struct {
	Value       NullableInteger `xml:",chardata"`
	Type        string          `xml:"type,attr" json:"type"`
	Unit        *string         `xml:"unit,attr" json:"unit,omitempty"`
	RefID       *uint8          `xml:"refID,attr" json:"refID,omitempty"`
	Condition   *string         `xml:"condition,attr" json:"condition,omitempty"`
	Description *string         `xml:"description,attr" json:"description,omitempty"`
}

type EbSnowfallDepth struct {
	Value       NullableFloat `xml:",chardata"`
	Type        string        `xml:"type,attr" json:"type"`
	Unit        *string       `xml:"unit,attr" json:"unit,omitempty"`
	RefID       *uint8        `xml:"refID,attr" json:"refID,omitempty"`
	Condition   *string       `xml:"condition,attr" json:"condition,omitempty"`
	Description *string       `xml:"description,attr" json:"description,omitempty"`
}

type EbSnowDepth struct {
	Value       NullableFloat `xml:",chardata"`
	Type        string        `xml:"type,attr" json:"type"`
	Unit        *string       `xml:"unit,attr" json:"unit,omitempty"`
	RefID       *uint8        `xml:"refID,attr" json:"refID,omitempty"`
	Condition   *string       `xml:"condition,attr" json:"condition,omitempty"`
	Description *string       `xml:"description,attr" json:"description,omitempty"`
}

type EbVisibility struct {
	Value       NullableFloat `xml:",chardata"`
	Type        string        `xml:"type,attr" json:"type"`
	Unit        *string       `xml:"unit,attr" json:"unit,omitempty"`
	RefID       *uint8        `xml:"refID,attr" json:"refID,omitempty"`
	Condition   *string       `xml:"condition,attr" json:"condition,omitempty"`
	Description *string       `xml:"description,attr" json:"description,omitempty"`
}

type EbWeather struct {
	Value       string  `xml:",chardata"`
	Type        string  `xml:"type,attr" json:"type"`
	RefID       *uint8  `xml:"refID,attr" json:"refID,omitempty"`
	Condition   *string `xml:"condition,attr" json:"condition,omitempty"`
	Description *string `xml:"description,attr" json:"description,omitempty"`
}

type EbWeatherCode struct {
	Value       NullableInteger `xml:",chardata"`
	Type        string          `xml:"type,attr" json:"type"`
	RefID       *uint8          `xml:"refID,attr" json:"refID,omitempty"`
	Condition   *string         `xml:"condition,attr" json:"condition,omitempty"`
	Description *string         `xml:"description,attr" json:"description,omitempty"`
}

type EbSynopsis struct {
	Value string `xml:",chardata"`
	Type  string `xml:"type,attr" json:"type"`
}

type EbWaveHeight struct {
	Value       NullableFloat `xml:",chardata"`
	Type        string        `xml:"type,attr" json:"type"`
	Unit        *string       `xml:"unit,attr" json:"unit,omitempty"`
	RefID       *uint8        `xml:"refID,attr" json:"refID,omitempty"`
	Condition   *string       `xml:"condition,attr" json:"condition,omitempty"`
	Description *string       `xml:"description,attr" json:"description,omitempty"`
}

type EbTidalLevel struct {
	Value       NullableFloat `xml:",chardata"`
	Type        string        `xml:"type,attr" json:"type"`
	Unit        *string       `xml:"unit,attr" json:"unit,omitempty"`
	RefID       *uint8        `xml:"refID,attr" json:"refID,omitempty"`
	Condition   *string       `xml:"condition,attr" json:"condition,omitempty"`
	Description *string       `xml:"description,attr" json:"description,omitempty"`
}

type EbSeaIce struct {
	Value       string  `xml:",chardata"`
	Type        string  `xml:"type,attr" json:"type"`
	Unit        *string `xml:"unit,attr" json:"unit,omitempty"`
	RefID       *uint8  `xml:"refID,attr" json:"refID,omitempty"`
	Condition   *string `xml:"condition,attr" json:"condition,omitempty"`
	Description *string `xml:"description,attr" json:"description,omitempty"`
}

type EbIcing struct {
	Value       string  `xml:",chardata"`
	Type        string  `xml:"type,attr" json:"type"`
	Unit        *string `xml:"unit,attr" json:"unit,omitempty"`
	RefID       *uint8  `xml:"refID,attr" json:"refID,omitempty"`
	Condition   *string `xml:"condition,attr" json:"condition,omitempty"`
	Description *string `xml:"description,attr" json:"description,omitempty"`
}

type EbReliabilityClass struct {
	Value     string  `xml:",chardata"`
	Type      string  `xml:"type,attr" json:"type"`
	RefID     *uint8  `xml:"refID,attr" json:"refID,omitempty"`
	Condition *string `xml:"condition,attr" json:"condition,omitempty"`
}

type EbReliabilityValue struct {
	Value     string  `xml:",chardata"`
	Type      string  `xml:"type,attr" json:"type"`
	RefID     *uint8  `xml:"refID,attr" json:"refID,omitempty"`
	Condition *string `xml:"condition,attr" json:"condition,omitempty"`
}

type EbPossibilityRankOfWarning struct {
	Value     string  `xml:",chardata"`
	Type      string  `xml:"type,attr" json:"type"`
	RefID     *uint8  `xml:"refID,attr" json:"refID,omitempty"`
	Condition *string `xml:"condition,attr" json:"condition,omitempty"`
}

type EbTyphoonClass struct {
	Value string `xml:",chardata"`
	Type  string `xml:"type,attr" json:"type"`
}

type EbAreaClass struct {
	Value string `xml:",chardata"`
	Type  string `xml:"type,attr" json:"type"`
}

type EbIntensityClass struct {
	Value string `xml:",chardata"`
	Type  string `xml:"type,attr" json:"type"`
}

type EbWaterLevel struct {
	Value       NullableFloat `xml:",chardata"`
	Type        string        `xml:"type,attr" json:"type"`
	Unit        *string       `xml:"unit,attr" json:"unit,omitempty"`
	RefID       *uint8        `xml:"refID,attr" json:"refID,omitempty"`
	Condition   *string       `xml:"condition,attr" json:"condition,omitempty"`
	Description *string       `xml:"description,attr" json:"description,omitempty"`
}

type EbFloodDepth struct {
	Value       NullableFloat `xml:",chardata"`
	Type        string        `xml:"type,attr" json:"type"`
	Unit        *string       `xml:"unit,attr" json:"unit,omitempty"`
	Bound       *string       `xml:"bound,attr" json:"bound,omitempty"`
	RefID       *uint8        `xml:"refID,attr" json:"refID,omitempty"`
	Condition   *string       `xml:"condition,attr" json:"condition,omitempty"`
	Description *string       `xml:"description,attr" json:"description,omitempty"`
}

type EbDischarge struct {
	Value       NullableFloat `xml:",chardata"`
	Type        string        `xml:"type,attr" json:"type"`
	Unit        *string       `xml:"unit,attr" json:"unit,omitempty"`
	RefID       *uint8        `xml:"refID,attr" json:"refID,omitempty"`
	Condition   *string       `xml:"condition,attr" json:"condition,omitempty"`
	Description *string       `xml:"description,attr" json:"description,omitempty"`
}

type EbClimateFeature struct {
	GeneralSituationTexts      []*EbReferableString           `xml:"GeneralSituationText" json:"generalSituationTexts,omitempty"`
	SignificantClimateElements []*EbSignificantClimateElement `xml:"SignificantClimateElement" json:"significantClimateElements,omitempty"`
}

type EbReferableString struct {
	Value string  `xml:",chardata"`
	Type  *string `xml:"type,attr" json:"type,omitempty"`
	RefID *uint8  `xml:"refID,attr" json:"refID,omitempty"`
}

type EbComparison struct {
	Value       NullableFloat `xml:",chardata"`
	Type        string        `xml:"type,attr" json:"type"`
	Unit        *string       `xml:"unit,attr" json:"unit,omitempty"`
	RefID       *uint8        `xml:"refID,attr" json:"refID,omitempty"`
	Condition   *string       `xml:"condition,attr" json:"condition,omitempty"`
	Description *string       `xml:"description,attr" json:"description,omitempty"`
}

type EbSignificantClimateElement struct {
	Kind                                  string               `xml:"kind,attr" json:"kind"`
	Texts                                 []*EbReferableString `xml:"Text" json:"texts"`
	ProbabilityOfBelowNormal              *EbProbability       `xml:"ProbabilityOfBelowNormal" json:"probabilityOfBelowNormal,omitempty"`
	ProbabilityOfNormal                   *EbProbability       `xml:"ProbabilityOfNormal" json:"probabilityOfNormal,omitempty"`
	ProbabilityOfAboveNormal              *EbProbability       `xml:"ProbabilityOfAboveNormal" json:"probabilityOfAboveNormal,omitempty"`
	ProbabilityOfSignificantlyBelowNormal *EbProbability       `xml:"ProbabilityOfSignificantlyBelowNormal" json:"probabilityOfSignificantlyBelowNormal,omitempty"`
	ProbabilityOfSignificantlyAboveNormal *EbProbability       `xml:"ProbabilityOfSignificantlyAboveNormal" json:"probabilityOfSignificantlyAboveNormal,omitempty"`
	ThresholdOfBelowNormal                *EbThreshold         `xml:"ThresholdOfBelowNormal" json:"thresholdOfBelowNormal,omitempty"`
	ThresholdOfAboveNormal                *EbThreshold         `xml:"ThresholdOfAboveNormal" json:"thresholdOfAboveNormal,omitempty"`
	ThresholdOfSignificantlyBelowNormal   *EbThreshold         `xml:"ThresholdOfSignificantlyBelowNormal" json:"thresholdOfSignificantlyBelowNormal,omitempty"`
	ThresholdOfSignificantlyAboveNormal   *EbThreshold         `xml:"ThresholdOfSignificantlyAboveNormal" json:"thresholdOfSignificantlyAboveNormal,omitempty"`
}

type EbClassThresholdOfAverage struct {
	ThresholdOfMinimum                  *EbThreshold `xml:"ThresholdOfMinimum" json:"thresholdOfMinimum,omitempty"`
	ThresholdOfSignificantlyBelowNormal *EbThreshold `xml:"ThresholdOfSignificantlyBelowNormal" json:"thresholdOfSignificantlyBelowNormal,omitempty"`
	ThresholdOfBelowNormal              *EbThreshold `xml:"ThresholdOfBelowNormal" json:"thresholdOfBelowNormal,omitempty"`
	ThresholdOfAboveNormal              *EbThreshold `xml:"ThresholdOfAboveNormal" json:"thresholdOfAboveNormal,omitempty"`
	ThresholdOfSignificantlyAboveNormal *EbThreshold `xml:"ThresholdOfSignificantlyAboveNormal" json:"thresholdOfSignificantlyAboveNormal,omitempty"`
	ThresholdOfMaximum                  *EbThreshold `xml:"ThresholdOfMaximum" json:"thresholdOfMaximum,omitempty"`
}

type EbProbability struct {
	Value       float64 `xml:",chardata"`
	Unit        *string `xml:"unit,attr" json:"unit,omitempty"`
	Bound       *string `xml:"bound,attr" json:"bound,omitempty"`
	Significant *bool   `xml:"significant,attr" json:"significant,omitempty"`
}

type EbThreshold struct {
	Value       float64 `xml:",chardata"`
	Type        *string `xml:"type,attr" json:"type,omitempty"`
	Unit        *string `xml:"unit,attr" json:"unit,omitempty"`
	Description *string `xml:"description,attr" json:"description,omitempty"`
	Bound       *string `xml:"bound,attr" json:"bound,omitempty"`
}

type EbClimateProbabilityValues struct {
	Kind                     string        `xml:"kind,attr" json:"kind"`
	RefID                    *uint8        `xml:"refID,attr" json:"refID,omitempty"`
	ProbabilityOfBelowNormal EbProbability `xml:"ProbabilityOfBelowNormal" json:"probabilityOfBelowNormal"`
	ProbabilityOfNormal      EbProbability `xml:"ProbabilityOfNormal" json:"probabilityOfNormal"`
	ProbabilityOfAboveNormal EbProbability `xml:"ProbabilityOfAboveNormal" json:"probabilityOfAboveNormal"`
}

type EbSolarZenithAngle struct {
	Value     float64 `xml:",chardata"`
	Unit      *string `xml:"unit,attr" json:"unit,omitempty"`
	RefID     *uint8  `xml:"refID,attr" json:"refID,omitempty"`
	Condition *string `xml:"condition,attr" json:"condition,omitempty"`
}

type EbUvIndex struct {
	Value       NullableFloat `xml:",chardata"`
	Type        *string       `xml:"type,attr" json:"type,omitempty"`
	RefID       *uint8        `xml:"refID,attr" json:"refID,omitempty"`
	Condition   *string       `xml:"condition,attr" json:"condition,omitempty"`
	Description *string       `xml:"description,attr" json:"description,omitempty"`
}

type EbSpeed struct {
	Value       NullableFloat `xml:",chardata"`
	Type        string        `xml:"type,attr" json:"type"`
	Unit        *string       `xml:"unit,attr" json:"unit,omitempty"`
	Condition   *string       `xml:"condition,attr" json:"condition,omitempty"`
	Description *string       `xml:"description,attr" json:"description,omitempty"`
}

type EbDirection struct {
	Value       string  `xml:",chardata"`
	Type        string  `xml:"type,attr" json:"type"`
	Unit        *string `xml:"unit,attr" json:"unit,omitempty"`
	Condition   *string `xml:"condition,attr" json:"condition,omitempty"`
	Description *string `xml:"description,attr" json:"description,omitempty"`
}

type EbBearings struct {
	Value       NullableInteger `xml:",chardata"`
	Type        string          `xml:"type,attr" json:"type"`
	Unit        *string         `xml:"unit,attr" json:"unit,omitempty"`
	Condition   *string         `xml:"condition,attr" json:"condition,omitempty"`
	Description *string         `xml:"description,attr" json:"description,omitempty"`
}

type EbRadius struct {
	Value       NullableFloat `xml:",chardata"`
	Type        string        `xml:"type,attr" json:"type"`
	Unit        *string       `xml:"unit,attr" json:"unit,omitempty"`
	Condition   *string       `xml:"condition,attr" json:"condition,omitempty"`
	Description *string       `xml:"description,attr" json:"description,omitempty"`
}

type EbMagnitude struct {
	Value       float64 `xml:",chardata"`
	Type        string  `xml:"type,attr" json:"type"`
	Condition   *string `xml:"condition,attr" json:"condition,omitempty"`
	Description *string `xml:"description,attr" json:"description,omitempty"`
}

type EbTsunamiHeight struct {
	Value       float64 `xml:",chardata"`
	Type        string  `xml:"type,attr" json:"type"`
	Unit        string  `xml:"unit,attr" json:"unit"`
	Condition   *string `xml:"condition,attr" json:"condition,omitempty"`
	Description *string `xml:"description,attr" json:"description,omitempty"`
}

type EbProbabilityOfAftershock struct {
	Value float64 `xml:",chardata"`
	Type  string  `xml:"type,attr" json:"type"`
	Unit  string  `xml:"unit,attr" json:"unit"`
}

type EbPlumeDirection struct {
	Value       string  `xml:",chardata"`
	Type        *string `xml:"type,attr" json:"type,omitempty"`
	Unit        *string `xml:"unit,attr" json:"unit,omitempty"`
	Condition   *string `xml:"condition,attr" json:"condition,omitempty"`
	Description *string `xml:"description,attr" json:"description,omitempty"`
}

type EbPlumeHeight struct {
	Value       NullableInteger `xml:",chardata"`
	Type        *string         `xml:"type,attr" json:"type,omitempty"`
	Unit        *string         `xml:"unit,attr" json:"unit,omitempty"`
	Condition   *string         `xml:"condition,attr" json:"condition,omitempty"`
	Description *string         `xml:"description,attr" json:"description,omitempty"`
}

type EbWeatherForecastProbability struct {
	Value float64 `xml:",chardata"`
	RefID *uint8  `xml:"refID,attr" json:"refID,omitempty"`
	Unit  *string `xml:"unit,attr" json:"unit,omitempty"`
}
