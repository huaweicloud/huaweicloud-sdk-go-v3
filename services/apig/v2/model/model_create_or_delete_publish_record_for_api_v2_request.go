package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateOrDeletePublishRecordForApiV2Request struct {
	// 实例编号

	InstanceId string `json:"instance_id"`

	Body *ApiPublishReq `json:"body,omitempty"`
}

func (o CreateOrDeletePublishRecordForApiV2Request) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateOrDeletePublishRecordForApiV2Request struct{}"
	}

	return strings.Join([]string{"CreateOrDeletePublishRecordForApiV2Request", string(data)}, " ")
}
