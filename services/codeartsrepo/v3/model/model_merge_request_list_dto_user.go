package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MergeRequestListDtoUser struct {

	// 当前用户是否可合入
	CanMerge *bool `json:"can_merge,omitempty"`
}

func (o MergeRequestListDtoUser) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MergeRequestListDtoUser struct{}"
	}

	return strings.Join([]string{"MergeRequestListDtoUser", string(data)}, " ")
}
