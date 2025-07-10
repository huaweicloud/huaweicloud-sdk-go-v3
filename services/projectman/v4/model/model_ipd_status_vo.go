package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IpdStatusVo 工作项状态查询接口返回状态数据
type IpdStatusVo struct {

	// 状态编码
	Code *string `json:"code,omitempty"`

	// 状态名称
	Name *string `json:"name,omitempty"`

	// 工作项的状态属性
	Belonging *string `json:"belonging,omitempty"`
}

func (o IpdStatusVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IpdStatusVo struct{}"
	}

	return strings.Join([]string{"IpdStatusVo", string(data)}, " ")
}
