package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ParamGroup struct {

	// 参数组ID。
	Id *string `json:"id,omitempty"`

	// 参数组名。
	Name *string `json:"name,omitempty"`
}

func (o ParamGroup) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ParamGroup struct{}"
	}

	return strings.Join([]string{"ParamGroup", string(data)}, " ")
}
