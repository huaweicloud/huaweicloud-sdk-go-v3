package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 批量重启边缘实例请求体。
type BatchRebootInstanceRequestBody struct {
	Reboot *BatchReboot `json:"reboot,omitempty"`
}

func (o BatchRebootInstanceRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchRebootInstanceRequestBody struct{}"
	}

	return strings.Join([]string{"BatchRebootInstanceRequestBody", string(data)}, " ")
}
