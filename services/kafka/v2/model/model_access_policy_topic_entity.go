package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 权限实体。
type AccessPolicyTopicEntity struct {

	// topic名称。
	Name string `json:"name" xml:"name"`

	// 权限列表。
	Policies []AccessPolicyEntity `json:"policies" xml:"policies"`
}

func (o AccessPolicyTopicEntity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AccessPolicyTopicEntity struct{}"
	}

	return strings.Join([]string{"AccessPolicyTopicEntity", string(data)}, " ")
}
