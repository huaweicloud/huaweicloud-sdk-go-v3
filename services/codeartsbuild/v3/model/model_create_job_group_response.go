package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateJobGroupResponse Response Object
type CreateJobGroupResponse struct {

	// 返回结果
	Result *[]JobGroupResponseBodyResult `json:"result,omitempty"`

	// 返回错误信息
	Error *string `json:"error,omitempty"`

	// 返回状态信息
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateJobGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateJobGroupResponse struct{}"
	}

	return strings.Join([]string{"CreateJobGroupResponse", string(data)}, " ")
}
