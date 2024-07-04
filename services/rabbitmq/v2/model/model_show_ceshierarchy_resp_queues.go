package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowCeshierarchyRespQueues struct {

	// Queue名称。
	Name *string `json:"name,omitempty"`

	// 对应的Vhost。
	Vhost *string `json:"vhost,omitempty"`
}

func (o ShowCeshierarchyRespQueues) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCeshierarchyRespQueues struct{}"
	}

	return strings.Join([]string{"ShowCeshierarchyRespQueues", string(data)}, " ")
}
