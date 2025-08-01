package common

import (
	"encoding/json"
)

func DecodeData(src, dst interface{}) error {
	byteSrc, err := json.Marshal(src)
	if err != nil {
		return err
	}

	err = json.Unmarshal(byteSrc, dst)
	if err != nil {
		return err
	}
	return nil
}
