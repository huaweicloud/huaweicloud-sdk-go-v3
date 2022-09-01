package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchRemovePeersFromChannelRequestBody struct {

	// 组织名称。仅可输入一个组织名称
	OrgName string `json:"org_name" xml:"org_name"`

	// 要退出的节点个数。取值范围(0, 组织中节点总数)
	Peers int32 `json:"peers" xml:"peers"`
}

func (o BatchRemovePeersFromChannelRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchRemovePeersFromChannelRequestBody struct{}"
	}

	return strings.Join([]string{"BatchRemovePeersFromChannelRequestBody", string(data)}, " ")
}
