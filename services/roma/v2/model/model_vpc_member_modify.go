package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VpcMemberModify struct {

	// 后端实例列表
	Members *[]MemberInfo `json:"members,omitempty"`

	// 需要修改的后端服务器组  不传时使用members中的定义对VPC通道后端进行全量覆盖修改。  传入时，只对members中对应后端服务器组的后端实例进行处理，其他后端服务器组的入参会被忽略。例如：member_group_name=primary时，只处理members中后端服务器组为primary的后端实例。
	MemberGroupName *string `json:"member_group_name,omitempty"`
}

func (o VpcMemberModify) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VpcMemberModify struct{}"
	}

	return strings.Join([]string{"VpcMemberModify", string(data)}, " ")
}
