package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RelationInfo struct {

	// |参数名称：关联资源ID。| |参数约束及描述：关联资源ID。|
	RelativeResourceId *string `json:"relative_resource_id,omitempty"`

	// |参数名称：关联资源类型。| |参数约束及描述：关联资源类型，父资源：PARENT；根资源：ROOT|
	RelativeType *string `json:"relative_type,omitempty"`
}

func (o RelationInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RelationInfo struct{}"
	}

	return strings.Join([]string{"RelationInfo", string(data)}, " ")
}
