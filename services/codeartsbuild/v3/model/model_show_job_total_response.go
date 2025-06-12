package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowJobTotalResponse Response Object
type ShowJobTotalResponse struct {
	Result *TotalResponseBodyResult `json:"result,omitempty"`

	// 状态信息
	Status *string `json:"status,omitempty"`

	// 错误
	Error          *interface{} `json:"error,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ShowJobTotalResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowJobTotalResponse struct{}"
	}

	return strings.Join([]string{"ShowJobTotalResponse", string(data)}, " ")
}
