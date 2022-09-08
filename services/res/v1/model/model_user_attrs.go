package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UserAttrs struct {

	// 数据类型。
	DataType *string `json:"data_type,omitempty"`

	// 用户名称。
	Name *string `json:"name,omitempty"`
}

func (o UserAttrs) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UserAttrs struct{}"
	}

	return strings.Join([]string{"UserAttrs", string(data)}, " ")
}
