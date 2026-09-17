package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EpsInfo EpsInfo对象
type EpsInfo struct {

	// 企业项目ID
	Id *string `json:"id,omitempty"`

	// 企业项目名称
	Name *string `json:"name,omitempty"`
}

func (o EpsInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EpsInfo struct{}"
	}

	return strings.Join([]string{"EpsInfo", string(data)}, " ")
}
