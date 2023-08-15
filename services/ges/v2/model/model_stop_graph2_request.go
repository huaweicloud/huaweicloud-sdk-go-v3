package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StopGraph2Request Request Object
type StopGraph2Request struct {

	// 图ID。
	GraphId string `json:"graph_id"`
}

func (o StopGraph2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopGraph2Request struct{}"
	}

	return strings.Join([]string{"StopGraph2Request", string(data)}, " ")
}
