package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResRelation struct {

	// |参数名称：当前费用对应的资源ID| |参数约束及描述：当前费用对应的资源ID|
	SelfResourceId *string `json:"self_resource_id,omitempty"`

	// |参数名称：当前费用对应资源ID关联的资源信息。| |参数约束及描述：当前费用对应资源ID关联的资源信息，数组个数不超过2层|
	RelationInfos *[]RelationInfo `json:"relation_infos,omitempty"`
}

func (o ResRelation) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResRelation struct{}"
	}

	return strings.Join([]string{"ResRelation", string(data)}, " ")
}
