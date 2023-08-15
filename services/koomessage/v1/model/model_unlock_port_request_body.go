package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UnlockPortRequestBody 通道号解绑服务号请求体。
type UnlockPortRequestBody struct {

	// 主键ID，取查询通道号绑定信息列表返回的ID字段。
	Id string `json:"id"`
}

func (o UnlockPortRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UnlockPortRequestBody struct{}"
	}

	return strings.Join([]string{"UnlockPortRequestBody", string(data)}, " ")
}
