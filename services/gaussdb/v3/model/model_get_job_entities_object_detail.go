package model

import (
	"encoding/json"

	"strings"
)

type GetJobEntitiesObjectDetail struct {
	Instance *GetJobEntitiesInstanceInfoDetail `json:"instance"`
	// 任务涉及到的资源ID。

	ResourceIds []string `json:"resource_ids"`
}

func (o GetJobEntitiesObjectDetail) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "GetJobEntitiesObjectDetail struct{}"
	}

	return strings.Join([]string{"GetJobEntitiesObjectDetail", string(data)}, " ")
}
