package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckPublishAppReq 校验发布应用名称请求结构体。
type CheckPublishAppReq struct {

	// 应用名称。
	Name string `json:"name"`
}

func (o CheckPublishAppReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckPublishAppReq struct{}"
	}

	return strings.Join([]string{"CheckPublishAppReq", string(data)}, " ")
}
