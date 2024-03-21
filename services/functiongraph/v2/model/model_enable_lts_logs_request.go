package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EnableLtsLogsRequest Request Object
type EnableLtsLogsRequest struct {

	// 消息体的类型（格式）
	ContentType string `json:"Content-Type"`
}

func (o EnableLtsLogsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnableLtsLogsRequest struct{}"
	}

	return strings.Join([]string{"EnableLtsLogsRequest", string(data)}, " ")
}
