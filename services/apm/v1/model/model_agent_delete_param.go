package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AgentDeleteParam 删除agent入参。
type AgentDeleteParam struct {

	// 实例列表。
	InstanceList []int64 `json:"instance_list"`

	// region英文名称。
	Region string `json:"region"`

	// 应用id。
	BusinessId int64 `json:"business_id"`
}

func (o AgentDeleteParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AgentDeleteParam struct{}"
	}

	return strings.Join([]string{"AgentDeleteParam", string(data)}, " ")
}
