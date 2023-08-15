package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DeregisterScenesReqScenes struct {

	// 场景名。
	Name *string `json:"name,omitempty"`

	// 要取消订阅的application 名字列表(当前不支持)。
	Applications *[]string `json:"applications,omitempty"`
}

func (o DeregisterScenesReqScenes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeregisterScenesReqScenes struct{}"
	}

	return strings.Join([]string{"DeregisterScenesReqScenes", string(data)}, " ")
}
