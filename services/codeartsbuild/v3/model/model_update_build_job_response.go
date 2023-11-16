package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateBuildJobResponse Response Object
type UpdateBuildJobResponse struct {
	Result *CreateBuildJobResponseBodyResult `json:"result,omitempty"`

	// 状态信息
	Status *string `json:"status,omitempty"`

	// 错误信息
	Error          *string `json:"error,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateBuildJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateBuildJobResponse struct{}"
	}

	return strings.Join([]string{"UpdateBuildJobResponse", string(data)}, " ")
}
