package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeregisterScenes2Request struct {

	// 图ID。
	GraphId string `json:"graph_id"`

	Body *DeregisterScenesReq `json:"body,omitempty"`
}

func (o DeregisterScenes2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeregisterScenes2Request struct{}"
	}

	return strings.Join([]string{"DeregisterScenes2Request", string(data)}, " ")
}
