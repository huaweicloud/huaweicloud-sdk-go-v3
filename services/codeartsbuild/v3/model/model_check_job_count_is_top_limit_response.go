package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckJobCountIsTopLimitResponse Response Object
type CheckJobCountIsTopLimitResponse struct {

	// 状态
	Status *string `json:"status,omitempty"`

	// 错误
	Error *interface{} `json:"error,omitempty"`

	// 结果
	Result         *bool `json:"result,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o CheckJobCountIsTopLimitResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckJobCountIsTopLimitResponse struct{}"
	}

	return strings.Join([]string{"CheckJobCountIsTopLimitResponse", string(data)}, " ")
}
