package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 批量关闭边缘实例请求体。
type BatchStopInstanceRequestBody struct {
	OsStop *BatchStop `json:"os-stop,omitempty" xml:"os-stop"`
}

func (o BatchStopInstanceRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchStopInstanceRequestBody struct{}"
	}

	return strings.Join([]string{"BatchStopInstanceRequestBody", string(data)}, " ")
}
