package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RestartGraph2Request Request Object
type RestartGraph2Request struct {

	// 图ID。
	GraphId string `json:"graph_id"`
}

func (o RestartGraph2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestartGraph2Request struct{}"
	}

	return strings.Join([]string{"RestartGraph2Request", string(data)}, " ")
}
