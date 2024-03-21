package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowProjectAsyncStatusLogInfoRequest Request Object
type ShowProjectAsyncStatusLogInfoRequest struct {

	// 消息体的类型（格式）
	ContentType string `json:"Content-Type"`
}

func (o ShowProjectAsyncStatusLogInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowProjectAsyncStatusLogInfoRequest struct{}"
	}

	return strings.Join([]string{"ShowProjectAsyncStatusLogInfoRequest", string(data)}, " ")
}
