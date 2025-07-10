package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowBestPracticeStatusResponse Response Object
type ShowBestPracticeStatusResponse struct {

	// 操作Id
	OperationId *string `json:"operation_id,omitempty"`

	// 状态：进行中，成功，成败
	Status *string `json:"status,omitempty"`

	// 检测进度
	PercentageComplete *int32 `json:"percentage_complete,omitempty"`

	// 错误消息
	ErrorMessages  *[]string `json:"error_messages,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ShowBestPracticeStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBestPracticeStatusResponse struct{}"
	}

	return strings.Join([]string{"ShowBestPracticeStatusResponse", string(data)}, " ")
}
