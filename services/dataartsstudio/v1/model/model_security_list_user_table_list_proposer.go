package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SecurityListUserTableListProposer 需要查询的人
type SecurityListUserTableListProposer struct {

	// 必填，用户、用户组id，iam接口获取
	Id string `json:"id"`

	// 非必填，用户、用户组name，iam接口获取
	Name *string `json:"name,omitempty"`

	// 必填，枚举，0用户
	Type int32 `json:"type"`
}

func (o SecurityListUserTableListProposer) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SecurityListUserTableListProposer struct{}"
	}

	return strings.Join([]string{"SecurityListUserTableListProposer", string(data)}, " ")
}
