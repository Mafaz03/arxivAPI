package timeinfo

import (
	"encoding/json"
	"errors"
	"os"
	"time"
)

type TimeInfo struct {
	StaticInfo        string    `json:"static_info"`
	LastRunTime       string    `json:"last_run_time"`
	LastRunTimeParsed time.Time `json:"-"`
}

func ReadData(filename string) (*TimeInfo, error) {
	encData, err := os.ReadFile(filename)
	if err != nil {
		return nil, errors.New("time related json file could not be read")
	}

	timeinfo := &TimeInfo{}
	err = json.Unmarshal(encData, timeinfo)
	if err != nil {
		return nil, errors.New("error unmarshaling json data")
	}

	if timeinfo.LastRunTime != "" {
		timeinfo.LastRunTimeParsed, err = time.Parse("2006-01-02 15:04:05", timeinfo.LastRunTime)
		if err != nil {
			return nil, errors.New("error parsing last run time")
		}
	}

	return timeinfo, nil
}

func WriteData(filename string, data *TimeInfo) error {

	encData, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		return err
	}

	err = os.WriteFile(filename, encData, 0644)
	if err != nil {
		return errors.New("time related json file could not be saved")
	}

	return nil
}

func UpdateLastRunTime(filename string) error {
	data, err := ReadData(filename)
	if err != nil {
		return err
	}

	data.LastRunTimeParsed = time.Now().UTC()
	data.LastRunTime = data.LastRunTimeParsed.Format("2006-01-02 15:04:05")
	err = WriteData(filename, data)
	if err != nil {
		return err
	}

	return nil
}
