package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DeleteInstancesSessionRequestBody struct {

	// 是否删除全部会话。
	IsAll bool `json:"is_all"`

	// 需要删除的会话id。is_all为false的时候，session_ids为必填，不能为空。
	SessionIds *[]string `json:"session_ids,omitempty"`
}

func (o DeleteInstancesSessionRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteInstancesSessionRequestBody struct{}"
	}

	return strings.Join([]string{"DeleteInstancesSessionRequestBody", string(data)}, " ")
}
