package fitbit

type SleepService service

type Sleep struct {
	Sleep []SleepDatapoint `json:"sleep"`
}

type SleepDatapoint struct {
	Data                string      `json:"dateOfSleep"`
	Duration            int         `json:"duration"`
	Efficiency          int         `json:"efficiency"`
	IsMainSleep         bool        `json:"isMainSleep"`
	Levels              SleepLevels `json:"levels"`
	LogId               int         `json:"logId"`
	MinutesAfterWakeup  int         `json:"minutesAfterWakeup"`
	MinutesAsleep       int         `json:"minutesAsleep"`
	MinutesAwake        int         `json:"minutesAwake"`
	MinutesToFallAsleep int         `json:"minutesToFallAsleep"`
	StartTime           string      `json:"startTime"`
	TimeInBed           int         `json:"timeInBed"`
	Type                string      `json:"type"`
}

type SleepLevels struct {
	Summary   SleepLevelsSummary     `json:"summary"`
	Data      []SleepLevelsDatapoint `json:"data"`
	ShortData []SleepLevelsDatapoint `json:"shortData"`
}

type SleepLevelsSummary struct {
	Deep  SleepLevel `json:"deep"`
	Light SleepLevel `json:"light"`
	Rem   SleepLevel `json:"rem"`
	Wake  SleepLevel `json:"wake"`
}
type SleepLevel struct {
	Count           int `json:"count"`
	Minutes         int `json:"minutes"`
	AvgMinutes30Day int `json:"thirtyDayAvgMinutes"`
}

type SleepLevelsDatapoint struct {
	Datetime string `json:"datetime"`
	Level    string `json:"level"`
	Duration int    `json:"seconds"`
}
