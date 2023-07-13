package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RegisterScenes2Request Request Object
type RegisterScenes2Request struct {

	// 图ID。
	GraphId string `json:"graph_id"`

	Body *RegisterScenesReq `json:"body,omitempty"`
}

func (o RegisterScenes2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RegisterScenes2Request struct{}"
	}

	return strings.Join([]string{"RegisterScenes2Request", string(data)}, " ")
}
