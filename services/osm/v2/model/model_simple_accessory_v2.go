package model

import (
	"encoding/json"

	"strings"
)

type SimpleAccessoryV2 struct {
	// 附件id

	AccessoryId *string `json:"accessory_id,omitempty"`
	// 附件实际名称

	FileActualName *string `json:"file_actual_name,omitempty"`
}

func (o SimpleAccessoryV2) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "SimpleAccessoryV2 struct{}"
	}

	return strings.Join([]string{"SimpleAccessoryV2", string(data)}, " ")
}
