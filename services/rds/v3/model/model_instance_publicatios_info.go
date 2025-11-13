package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InstancePublicatiosInfo 实例发布信息。
type InstancePublicatiosInfo struct {

	// 实例id。
	InstanceId *string `json:"instance_id,omitempty"`

	// 实例名称。
	InstanceName *string `json:"instance_name,omitempty"`

	// 发布id。
	PublicationId *string `json:"publication_id,omitempty"`

	// 发布名称。
	PublicationName *string `json:"publication_name,omitempty"`
}

func (o InstancePublicatiosInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstancePublicatiosInfo struct{}"
	}

	return strings.Join([]string{"InstancePublicatiosInfo", string(data)}, " ")
}
