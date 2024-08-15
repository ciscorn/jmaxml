// Code generated by jmx_codegen; DO NOT EDIT.

package jmaxml

import "time"

type SeisBody struct {
	Naming          *SeisNaming          `xml:"Naming" json:"naming,omitempty"`
	Tsunami         *SeisTsunami         `xml:"Tsunami" json:"tsunami,omitempty"`
	Earthquakes     []*SeisEarthquake    `xml:"Earthquake" json:"earthquakes,omitempty"`
	Intensity       *SeisIntensity       `xml:"Intensity" json:"intensity,omitempty"`
	Tokai           *SeisTokai           `xml:"Tokai" json:"tokai,omitempty"`
	EarthquakeInfo  *SeisEarthquakeInfo  `xml:"EarthquakeInfo" json:"earthquakeInfo,omitempty"`
	EarthquakeCount *SeisEarthquakeCount `xml:"EarthquakeCount" json:"earthquakeCount,omitempty"`
	Aftershock      *SeisAftershocks     `xml:"Aftershock" json:"aftershock,omitempty"`
	Text            *string              `xml:"Text" json:"text,omitempty"`
	NextAdvisory    *string              `xml:"NextAdvisory" json:"nextAdvisory,omitempty"`
	Comments        *SeisComment         `xml:"Comments" json:"comments,omitempty"`
}

type SeisEarthquake struct {
	OriginTime  *time.Time      `xml:"OriginTime" json:"originTime,omitempty"`
	ArrivalTime time.Time       `xml:"ArrivalTime" json:"arrivalTime"`
	Condition   *string         `xml:"Condition" json:"condition,omitempty"`
	Hypocenter  *SeisHypocenter `xml:"Hypocenter" json:"hypocenter,omitempty"`
	Magnitudes  []*EbMagnitude  `xml:"Magnitude" json:"magnitudes"`
}

type SeisHypocenter struct {
	Area     SeisHypoArea  `xml:"Area" json:"area"`
	Source   *string       `xml:"Source" json:"source,omitempty"`
	Accuracy *SeisAccuracy `xml:"Accuracy" json:"accuracy,omitempty"`
}

type SeisHypoArea struct {
	Name         string                    `xml:"Name" json:"name"`
	Code         SeisHypoAreaCode          `xml:"Code" json:"code"`
	Coordinates  []*EbCoordinate           `xml:"Coordinate" json:"coordinates"`
	ReduceName   *string                   `xml:"ReduceName" json:"reduceName,omitempty"`
	ReduceCode   *SeisHypoAreaReduceCode   `xml:"ReduceCode" json:"reduceCode,omitempty"`
	DetailedName *string                   `xml:"DetailedName" json:"detailedName,omitempty"`
	DetailedCode *SeisHypoAreaDetailedCode `xml:"DetailedCode" json:"detailedCode,omitempty"`
	NameFromMark *string                   `xml:"NameFromMark" json:"nameFromMark,omitempty"`
	MarkCode     *SeisHypoAreaMarkCode     `xml:"MarkCode" json:"markCode,omitempty"`
	Direction    *string                   `xml:"Direction" json:"direction,omitempty"`
	Distance     *SeisHypoAreaDistance     `xml:"Distance" json:"distance,omitempty"`
	LandOrSea    *string                   `xml:"LandOrSea" json:"landOrSea,omitempty"`
}

type SeisHypoAreaCode struct {
	Value string `xml:",chardata"`
	Type  string `xml:"type,attr" json:"type"`
}

type SeisHypoAreaReduceCode struct {
	Value string `xml:",chardata"`
	Type  string `xml:"type,attr" json:"type"`
}

type SeisHypoAreaDetailedCode struct {
	Value string `xml:",chardata"`
	Type  string `xml:"type,attr" json:"type"`
}

type SeisHypoAreaMarkCode struct {
	Value string `xml:",chardata"`
	Type  string `xml:"type,attr" json:"type"`
}

type SeisHypoAreaDistance struct {
	Value int32  `xml:",chardata"`
	Unit  string `xml:"unit,attr" json:"unit"`
}

type SeisAccuracy struct {
	Epicenter                    SeisAccuracyEpicenter `xml:"Epicenter" json:"epicenter"`
	Depth                        SeisAccuracyDepth     `xml:"Depth" json:"depth"`
	MagnitudeCalculation         SeisAccuracyMagnitude `xml:"MagnitudeCalculation" json:"magnitudeCalculation"`
	NumberOfMagnitudeCalculation int32                 `xml:"NumberOfMagnitudeCalculation" json:"numberOfMagnitudeCalculation"`
}

type SeisAccuracyEpicenter struct {
	Value float64 `xml:",chardata"`
	Rank  int32   `xml:"rank,attr" json:"rank"`
	Rank2 int32   `xml:"rank2,attr" json:"rank2"`
}

type SeisAccuracyDepth struct {
	Value float64 `xml:",chardata"`
	Rank  int32   `xml:"rank,attr" json:"rank"`
}

type SeisAccuracyMagnitude struct {
	Value float64 `xml:",chardata"`
	Rank  int32   `xml:"rank,attr" json:"rank"`
}

type SeisTsunami struct {
	Release     *string            `xml:"Release" json:"release,omitempty"`
	Observation *SeisTsunamiDetail `xml:"Observation" json:"observation,omitempty"`
	Estimation  *SeisTsunamiDetail `xml:"Estimation" json:"estimation,omitempty"`
	Forecast    *SeisTsunamiDetail `xml:"Forecast" json:"forecast,omitempty"`
}

type SeisTsunamiDetail struct {
	CodeDefine *SeisCodeDefine    `xml:"CodeDefine" json:"codeDefine,omitempty"`
	Items      []*SeisTsunamiItem `xml:"Item" json:"items"`
}

type SeisTsunamiItem struct {
	Area        SeisForecastArea      `xml:"Area" json:"area"`
	Category    *SeisCategory         `xml:"Category" json:"category,omitempty"`
	FirstHeight *SeisFirstHeight      `xml:"FirstHeight" json:"firstHeight,omitempty"`
	MaxHeight   *SeisMaxHeight        `xml:"MaxHeight" json:"maxHeight,omitempty"`
	Duration    *Duration             `xml:"Duration" json:"duration,omitempty"`
	Stations    []*SeisTsunamiStation `xml:"Station" json:"stations,omitempty"`
}

type SeisForecastArea struct {
	Name   string              `xml:"Name" json:"name"`
	Code   string              `xml:"Code" json:"code"`
	Cities []*SeisForecastCity `xml:"City" json:"cities,omitempty"`
}

type SeisForecastCity struct {
	Name string `xml:"Name" json:"name"`
	Code string `xml:"Code" json:"code"`
}

type SeisCategory struct {
	Kind     SeisKind  `xml:"Kind" json:"kind"`
	LastKind *SeisKind `xml:"LastKind" json:"lastKind,omitempty"`
}

type SeisKind struct {
	Name string `xml:"Name" json:"name"`
	Code string `xml:"Code" json:"code"`
}

type SeisFirstHeight struct {
	ArrivalTimeFrom *time.Time       `xml:"ArrivalTimeFrom" json:"arrivalTimeFrom,omitempty"`
	ArrivalTimeTo   *time.Time       `xml:"ArrivalTimeTo" json:"arrivalTimeTo,omitempty"`
	ArrivalTime     *time.Time       `xml:"ArrivalTime" json:"arrivalTime,omitempty"`
	Condition       *string          `xml:"Condition" json:"condition,omitempty"`
	Initial         *string          `xml:"Initial" json:"initial,omitempty"`
	TsunamiHeight   *EbTsunamiHeight `xml:"TsunamiHeight" json:"tsunamiHeight,omitempty"`
	Revise          *string          `xml:"Revise" json:"revise,omitempty"`
	Period          *float64         `xml:"Period" json:"period,omitempty"`
}

type SeisMaxHeight struct {
	DateTime          *time.Time       `xml:"DateTime" json:"dateTime,omitempty"`
	Condition         *string          `xml:"Condition" json:"condition,omitempty"`
	TsunamiHeightFrom *EbTsunamiHeight `xml:"TsunamiHeightFrom" json:"tsunamiHeightFrom,omitempty"`
	TsunamiHeightTo   *EbTsunamiHeight `xml:"TsunamiHeightTo" json:"tsunamiHeightTo,omitempty"`
	TsunamiHeight     *EbTsunamiHeight `xml:"TsunamiHeight" json:"tsunamiHeight,omitempty"`
	Revise            *string          `xml:"Revise" json:"revise,omitempty"`
	Period            *float64         `xml:"Period" json:"period,omitempty"`
}

type SeisCurrentHeight struct {
	StartTime     *time.Time       `xml:"StartTime" json:"startTime,omitempty"`
	EndTime       *time.Time       `xml:"EndTime" json:"endTime,omitempty"`
	Condition     *string          `xml:"Condition" json:"condition,omitempty"`
	TsunamiHeight *EbTsunamiHeight `xml:"TsunamiHeight" json:"tsunamiHeight,omitempty"`
}

type SeisTsunamiStation struct {
	Name             string             `xml:"Name" json:"name"`
	Code             string             `xml:"Code" json:"code"`
	Sensor           *string            `xml:"Sensor" json:"sensor,omitempty"`
	HighTideDateTime *time.Time         `xml:"HighTideDateTime" json:"highTideDateTime,omitempty"`
	FirstHeight      SeisFirstHeight    `xml:"FirstHeight" json:"firstHeight"`
	MaxHeight        *SeisMaxHeight     `xml:"MaxHeight" json:"maxHeight,omitempty"`
	CurrentHeight    *SeisCurrentHeight `xml:"CurrentHeight" json:"currentHeight,omitempty"`
}

type SeisIntensity struct {
	Forecast    *SeisIntensityDetail `xml:"Forecast" json:"forecast,omitempty"`
	Observation *SeisIntensityDetail `xml:"Observation" json:"observation,omitempty"`
}

type SeisIntensityDetail struct {
	CodeDefine    *SeisCodeDefine        `xml:"CodeDefine" json:"codeDefine,omitempty"`
	MaxInt        *string                `xml:"MaxInt" json:"maxInt,omitempty"`
	MaxLgInt      *string                `xml:"MaxLgInt" json:"maxLgInt,omitempty"`
	LgCategory    *string                `xml:"LgCategory" json:"lgCategory,omitempty"`
	ForecastInt   *SeisForecastInt       `xml:"ForecastInt" json:"forecastInt,omitempty"`
	ForecastLgInt *SeisForecastLgInt     `xml:"ForecastLgInt" json:"forecastLgInt,omitempty"`
	Appendix      *SeisIntensityAppendix `xml:"Appendix" json:"appendix,omitempty"`
	Prefs         []*SeisIntensityPref   `xml:"Pref" json:"prefs,omitempty"`
}

type SeisForecastInt struct {
	Bound *string `xml:"bound,attr" json:"bound,omitempty"`
	From  *string `xml:"From" json:"from,omitempty"`
	To    *string `xml:"To" json:"to,omitempty"`
}

type SeisForecastLgInt struct {
	Bound *string `xml:"bound,attr" json:"bound,omitempty"`
	From  *string `xml:"From" json:"from,omitempty"`
	To    *string `xml:"To" json:"to,omitempty"`
}

type SeisIntensityAppendix struct {
	MaxIntChange         int32  `xml:"MaxIntChange" json:"maxIntChange"`
	MaxLgIntChange       *int32 `xml:"MaxLgIntChange" json:"maxLgIntChange,omitempty"`
	MaxIntChangeReason   int32  `xml:"MaxIntChangeReason" json:"maxIntChangeReason"`
	MaxLgIntChangeReason *int32 `xml:"MaxLgIntChangeReason" json:"maxLgIntChangeReason,omitempty"`
}

type SeisIntensityPref struct {
	Name          string               `xml:"Name" json:"name"`
	Code          string               `xml:"Code" json:"code"`
	Category      *SeisCategory        `xml:"Category" json:"category,omitempty"`
	MaxInt        *string              `xml:"MaxInt" json:"maxInt,omitempty"`
	MaxLgInt      *string              `xml:"MaxLgInt" json:"maxLgInt,omitempty"`
	ForecastInt   *SeisForecastInt     `xml:"ForecastInt" json:"forecastInt,omitempty"`
	ForecastLgInt *SeisForecastLgInt   `xml:"ForecastLgInt" json:"forecastLgInt,omitempty"`
	ArrivalTime   *time.Time           `xml:"ArrivalTime" json:"arrivalTime,omitempty"`
	Condition     *string              `xml:"Condition" json:"condition,omitempty"`
	Revise        *string              `xml:"Revise" json:"revise,omitempty"`
	Areas         []*SeisIntensityArea `xml:"Area" json:"areas,omitempty"`
}

type SeisIntensityArea struct {
	Name              string                  `xml:"Name" json:"name"`
	Code              string                  `xml:"Code" json:"code"`
	Category          *SeisCategory           `xml:"Category" json:"category,omitempty"`
	MaxInt            *string                 `xml:"MaxInt" json:"maxInt,omitempty"`
	MaxLgInt          *string                 `xml:"MaxLgInt" json:"maxLgInt,omitempty"`
	ForecastInt       *SeisForecastInt        `xml:"ForecastInt" json:"forecastInt,omitempty"`
	ForecastLgInt     *SeisForecastLgInt      `xml:"ForecastLgInt" json:"forecastLgInt,omitempty"`
	ArrivalTime       *time.Time              `xml:"ArrivalTime" json:"arrivalTime,omitempty"`
	Condition         *string                 `xml:"Condition" json:"condition,omitempty"`
	Revise            *string                 `xml:"Revise" json:"revise,omitempty"`
	Cities            []*SeisIntensityCity    `xml:"City" json:"cities,omitempty"`
	IntensityStations []*SeisIntensityStation `xml:"IntensityStation" json:"intensityStations,omitempty"`
}

type SeisIntensityCity struct {
	Name              string                  `xml:"Name" json:"name"`
	Code              string                  `xml:"Code" json:"code"`
	Category          *SeisCategory           `xml:"Category" json:"category,omitempty"`
	MaxInt            *string                 `xml:"MaxInt" json:"maxInt,omitempty"`
	MaxLgInt          *string                 `xml:"MaxLgInt" json:"maxLgInt,omitempty"`
	ForecastInt       *SeisForecastInt        `xml:"ForecastInt" json:"forecastInt,omitempty"`
	ForecastLgInt     *SeisForecastLgInt      `xml:"ForecastLgInt" json:"forecastLgInt,omitempty"`
	ArrivalTime       *time.Time              `xml:"ArrivalTime" json:"arrivalTime,omitempty"`
	Condition         *string                 `xml:"Condition" json:"condition,omitempty"`
	Revise            *string                 `xml:"Revise" json:"revise,omitempty"`
	IntensityStations []*SeisIntensityStation `xml:"IntensityStation" json:"intensityStations,omitempty"`
}

type SeisIntensityStation struct {
	Name            string                `xml:"Name" json:"name"`
	Code            string                `xml:"Code" json:"code"`
	Int             *string               `xml:"Int" json:"int,omitempty"`
	K               *float64              `xml:"K" json:"k,omitempty"`
	LgInt           *string               `xml:"LgInt" json:"lgInt,omitempty"`
	LgIntPerPeriods []*SeisLgIntPerPeriod `xml:"LgIntPerPeriod" json:"lgIntPerPeriods,omitempty"`
	Sva             *SeisSva              `xml:"Sva" json:"sva,omitempty"`
	SvaPerPeriods   []*SeisSvaPerPeriod   `xml:"SvaPerPeriod" json:"svaPerPeriods,omitempty"`
	Revise          *string               `xml:"Revise" json:"revise,omitempty"`
}

type SeisLgIntPerPeriod struct {
	Value        string   `xml:",chardata"`
	PeriodicBand *int32   `xml:"PeriodicBand,attr" json:"PeriodicBand,omitempty"`
	Period       *float64 `xml:"Period,attr" json:"Period,omitempty"`
	PeriodUnit   *string  `xml:"PeriodUnit,attr" json:"PeriodUnit,omitempty"`
}

type SeisSva struct {
	Value float64 `xml:",chardata"`
	Unit  string  `xml:"unit,attr" json:"unit"`
}

type SeisSvaPerPeriod struct {
	Value        float64  `xml:",chardata"`
	Unit         string   `xml:"unit,attr" json:"unit"`
	PeriodicBand *int32   `xml:"PeriodicBand,attr" json:"PeriodicBand,omitempty"`
	Period       *float64 `xml:"Period,attr" json:"Period,omitempty"`
	PeriodUnit   *string  `xml:"PeriodUnit,attr" json:"PeriodUnit,omitempty"`
}

type SeisEarthquakeCount struct {
	Items []*SeisCountData `xml:"Item" json:"items"`
}

type SeisCountData struct {
	Type       string    `xml:"type,attr" json:"type"`
	StartTime  time.Time `xml:"StartTime" json:"startTime"`
	EndTime    time.Time `xml:"EndTime" json:"endTime"`
	Number     int32     `xml:"Number" json:"number"`
	FeltNumber int32     `xml:"FeltNumber" json:"feltNumber"`
	Condition  *string   `xml:"Condition" json:"condition,omitempty"`
}

type SeisTokai struct {
	InfoKind   string          `xml:"InfoKind" json:"infoKind"`
	InfoSerial *SeisInfoSerial `xml:"InfoSerial" json:"infoSerial,omitempty"`
	Text       string          `xml:"Text" json:"text"`
}

type SeisInfoSerial struct {
	CodeType string `xml:"codeType,attr" json:"codeType"`
	Name     string `xml:"Name" json:"name"`
	Code     string `xml:"Code" json:"code"`
}

type SeisEarthquakeInfo struct {
	Type       string          `xml:"type,attr" json:"type"`
	InfoKind   string          `xml:"InfoKind" json:"infoKind"`
	InfoSerial *SeisInfoSerial `xml:"InfoSerial" json:"infoSerial,omitempty"`
	Text       string          `xml:"Text" json:"text"`
	Appendix   *string         `xml:"Appendix" json:"appendix,omitempty"`
}

type SeisNaming struct {
	Value   string  `xml:",chardata"`
	English *string `xml:"english,attr" json:"english,omitempty"`
}

type SeisAftershocks struct {
	Items []*SeisAftershockItem `xml:"Item" json:"items"`
	Text  *string               `xml:"Text" json:"text,omitempty"`
}

type SeisAftershockItem struct {
	StartTime               time.Time                 `xml:"StartTime" json:"startTime"`
	EndTime                 time.Time                 `xml:"EndTime" json:"endTime"`
	ProbabilityOfAftershock EbProbabilityOfAftershock `xml:"ProbabilityOfAftershock" json:"probabilityOfAftershock"`
	TargetMagnitude         EbMagnitude               `xml:"TargetMagnitude" json:"targetMagnitude"`
	Text                    *string                   `xml:"Text" json:"text,omitempty"`
}

type SeisComment struct {
	WarningComment     *SeisCommentForm `xml:"WarningComment" json:"warningComment,omitempty"`
	ForecastComment    *SeisCommentForm `xml:"ForecastComment" json:"forecastComment,omitempty"`
	ObservationComment *SeisCommentForm `xml:"ObservationComment" json:"observationComment,omitempty"`
	VarComment         *SeisCommentForm `xml:"VarComment" json:"varComment,omitempty"`
	FreeFormComment    *string          `xml:"FreeFormComment" json:"freeFormComment,omitempty"`
	URI                *string          `xml:"URI" json:"uRI,omitempty"`
}

type SeisCommentForm struct {
	CodeType string     `xml:"codeType,attr" json:"codeType"`
	Text     string     `xml:"Text" json:"text"`
	Codes    StringList `xml:"Code" json:"codes"`
}

type SeisCodeDefine struct {
	Types []*SeisCodeDefineType `xml:"Type" json:"types"`
}

type SeisCodeDefineType struct {
	Value string `xml:",chardata"`
	Xpath string `xml:"xpath,attr" json:"xpath"`
}
