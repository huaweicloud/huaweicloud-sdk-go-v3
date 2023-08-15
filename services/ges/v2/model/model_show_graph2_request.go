package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowGraph2Request Request Object
type ShowGraph2Request struct {

	// 图ID。
	GraphId string `json:"graph_id"`
}

func (o ShowGraph2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowGraph2Request struct{}"
	}

	return strings.Join([]string{"ShowGraph2Request", string(data)}, " ")
}
