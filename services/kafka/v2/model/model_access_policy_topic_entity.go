package model

import (
	"encoding/json"

	"strings"
)

// 权限实体。
type AccessPolicyTopicEntity struct {
	// topic名称。

	Name string `json:"name"`
	// 权限列表。

	Policies []AccessPolicyEntity `json:"policies"`
}

func (o AccessPolicyTopicEntity) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "AccessPolicyTopicEntity struct{}"
	}

	return strings.Join([]string{"AccessPolicyTopicEntity", string(data)}, " ")
}
