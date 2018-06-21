package fitbit

type HeartRateService service
type HeartRate int64
type Calorie float64

type Heart struct {
	ActivitiesHeart         []ActivityHeart         `json:"activities-heart"`
	ActivitiesHeartIntraday []ActivityHeartIntraday `json:"activities-heart-intraday"`
}

type ActivityHeart struct {
	DateTime string             `json:"dateTime"`
	Value    ActivityHeartValue `json:"value"`
}

type ActivityHeartValue struct {
	CustomHeartRateZones []CustomHeartRateZone `json:"customHeartRateZones"`
	HeartRateZones       []HeartRateZone       `json:"heartRateZones"`
	RestingHeartRate     HeartRate             `json:"restingHeartRate"`
}

type CustomHeartRateZone struct {
	// TODO
}

type HeartRateZone struct {
	CaloriesOut Calorie   `json:"caloriesOut"`
	Max         HeartRate `json:"max"`
	Min         HeartRate `json:"min"`
	Minutes     HeartRate `json:"minutes"`
	Name        string    `json:"name"`
}

type ActivityHeartIntraday struct {
	Dataset  []HeartIntradayDatapoint `json:"dataset"`
	Interval int                      `json:"datasetInterval"`
	Type     string                   `json:"datasetType"`
}

type HeartIntradayDatapoint struct {
	Time      string    `json:"time"`
	HeartRate HeartRate `json:"value"`
}
