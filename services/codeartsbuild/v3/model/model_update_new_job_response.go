package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateNewJobResponse Response Object
type UpdateNewJobResponse struct {
	Result *CreateBuildJobResponseBodyResult `json:"result,omitempty"`

	// 状态信息
	Status *string `json:"status,omitempty"`

	// 错误信息
	Error          *string `json:"error,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateNewJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateNewJobResponse struct{}"
	}

	return strings.Join([]string{"UpdateNewJobResponse", string(data)}, " ")
}
