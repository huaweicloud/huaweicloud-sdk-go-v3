package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FailureForBatchCreateGroupMembersDto 关联失败成员
type FailureForBatchCreateGroupMembersDto struct {

	// iam_id
	IamId *string `json:"iam_id,omitempty"`

	// 失败原因
	Message *[]string `json:"message,omitempty"`
}

func (o FailureForBatchCreateGroupMembersDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FailureForBatchCreateGroupMembersDto struct{}"
	}

	return strings.Join([]string{"FailureForBatchCreateGroupMembersDto", string(data)}, " ")
}
