package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RegisterScenes请求体
type RegisterScenesReq struct {

	// 要订阅的具体场景。
	Scenes *[]RegisterScenesReqScenes `json:"scenes,omitempty"`
}

func (o RegisterScenesReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RegisterScenesReq struct{}"
	}

	return strings.Join([]string{"RegisterScenesReq", string(data)}, " ")
}
