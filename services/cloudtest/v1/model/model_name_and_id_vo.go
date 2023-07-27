package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NameAndIdVo 服务类型信息
type NameAndIdVo struct {

	// ID
	Id *string `json:"id,omitempty"`

	// 名称
	Name *string `json:"name,omitempty"`
}

func (o NameAndIdVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NameAndIdVo struct{}"
	}

	return strings.Join([]string{"NameAndIdVo", string(data)}, " ")
}
