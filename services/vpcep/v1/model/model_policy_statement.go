package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PolicyStatement policy
type PolicyStatement struct {

	// 允许或拒绝，控制访问权限
	Effect string `json:"Effect"`

	// obs访问权限
	Action []string `json:"Action"`

	// obs对象
	Resource []string `json:"Resource"`
}

func (o PolicyStatement) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PolicyStatement struct{}"
	}

	return strings.Join([]string{"PolicyStatement", string(data)}, " ")
}
