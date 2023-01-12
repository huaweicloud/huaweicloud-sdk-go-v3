package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowJob2Request struct {

	// 图ID。
	GraphId string `json:"graph_id"`

	// Job ID。
	JobId string `json:"job_id"`
}

func (o ShowJob2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowJob2Request struct{}"
	}

	return strings.Join([]string{"ShowJob2Request", string(data)}, " ")
}
