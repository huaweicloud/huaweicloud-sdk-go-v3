package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateBuildJobResponse Response Object
type CreateBuildJobResponse struct {
	Result *CreateBuildJobResponseBodyResult `json:"result,omitempty"`

	// 状态信息
	Status *string `json:"status,omitempty"`

	// 错误信息
	Error          *string `json:"error,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateBuildJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateBuildJobResponse struct{}"
	}

	return strings.Join([]string{"CreateBuildJobResponse", string(data)}, " ")
}
