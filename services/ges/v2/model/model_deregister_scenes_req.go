package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeregisterScenesReq DeregisterScenes请求体
type DeregisterScenesReq struct {

	// 要取消订阅的具体场景列表。
	Scenes *[]DeregisterScenesReqScenes `json:"scenes,omitempty"`
}

func (o DeregisterScenesReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeregisterScenesReq struct{}"
	}

	return strings.Join([]string{"DeregisterScenesReq", string(data)}, " ")
}
