package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UnsubscribeParam 资源取消或退订请求体。注：当前仅支持按需场景资源取消。
type UnsubscribeParam struct {

	// 资源id列表
	ResourceIds *[]string `json:"resource_ids,omitempty"`
}

func (o UnsubscribeParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UnsubscribeParam struct{}"
	}

	return strings.Join([]string{"UnsubscribeParam", string(data)}, " ")
}
