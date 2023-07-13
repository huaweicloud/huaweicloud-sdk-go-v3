package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StatusVo 状态
type StatusVo struct {

	// 状态id
	Id *string `json:"id,omitempty"`

	// 状态id对应的值
	Name *string `json:"name,omitempty"`
}

func (o StatusVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StatusVo struct{}"
	}

	return strings.Join([]string{"StatusVo", string(data)}, " ")
}
