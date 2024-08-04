package timeinfo

import (
	"encoding/json"
	"errors"
	"os"
)

type timeInfo struct {
	Static_info string `json:"static_info"`
	Last_run_time string `json:"last_run_time"`
}

func readData(filename string) (*timeInfo, error) {
    encData, err := os.ReadFile(filename)
    if err != nil {
        return nil, errors.New("time related json file could not be read")
    }

    timeinfo := &timeInfo{}
    err = json.Unmarshal(encData, timeinfo)
    if err != nil {
        return nil, errors.New("error unmarshaling json data")
    }

    return timeinfo, nil
}

func writeData(filename string, data *timeInfo) error {

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