package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SimpleUserIn struct {

	// 用户id
	Id *int32 `json:"id,omitempty"`

	// 用户uuid
	Identifier *string `json:"identifier,omitempty"`

	// 用户名称
	Name *string `json:"name,omitempty"`
}

func (o SimpleUserIn) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SimpleUserIn struct{}"
	}

	return strings.Join([]string{"SimpleUserIn", string(data)}, " ")
}
