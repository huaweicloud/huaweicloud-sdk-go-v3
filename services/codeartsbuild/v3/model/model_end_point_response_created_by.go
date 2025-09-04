package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EndPointResponseCreatedBy 创建者
type EndPointResponseCreatedBy struct {

	// 创建者用户名
	Username *string `json:"username,omitempty"`

	// 创建者用户id
	UserId *string `json:"user_id,omitempty"`
}

func (o EndPointResponseCreatedBy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EndPointResponseCreatedBy struct{}"
	}

	return strings.Join([]string{"EndPointResponseCreatedBy", string(data)}, " ")
}
