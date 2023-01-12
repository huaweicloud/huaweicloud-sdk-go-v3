package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RegisterScenesReqScenes struct {

	// 场景名称。
	Name *string `json:"name,omitempty"`

	// 要订阅的application名字列表(当前不支持)。
	Applications *[]string `json:"applications,omitempty"`
}

func (o RegisterScenesReqScenes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RegisterScenesReqScenes struct{}"
	}

	return strings.Join([]string{"RegisterScenesReqScenes", string(data)}, " ")
}
