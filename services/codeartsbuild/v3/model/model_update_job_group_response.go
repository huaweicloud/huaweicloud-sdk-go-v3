package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateJobGroupResponse Response Object
type UpdateJobGroupResponse struct {

	// 返回结果
	Result *[]JobGroupResponseBodyResult `json:"result,omitempty"`

	// 返回错误信息
	Error *string `json:"error,omitempty"`

	// 返回状态信息
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateJobGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateJobGroupResponse struct{}"
	}

	return strings.Join([]string{"UpdateJobGroupResponse", string(data)}, " ")
}
