package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteBuildJobsResponse Response Object
type BatchDeleteBuildJobsResponse struct {
	Result *DeleteTheJobResponseBodyResult `json:"result,omitempty"`

	// 状态信息
	Status *string `json:"status,omitempty"`

	// 返回错误信息
	Error          *string `json:"error,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o BatchDeleteBuildJobsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteBuildJobsResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteBuildJobsResponse", string(data)}, " ")
}
