package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateNewJobResponse Response Object
type CreateNewJobResponse struct {
	Result *CreateBuildJobResponseBodyResult `json:"result,omitempty"`

	// 状态信息
	Status *string `json:"status,omitempty"`

	// 错误信息
	Error          *string `json:"error,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateNewJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateNewJobResponse struct{}"
	}

	return strings.Join([]string{"CreateNewJobResponse", string(data)}, " ")
}
