package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OrderCatalogue 服务单目录-后续结合项目信息
type OrderCatalogue struct {

	// 服务单类型编码
	Code *string `json:"code,omitempty"`

	// 类型
	Type *string `json:"type,omitempty"`

	// 子类型
	SubType *string `json:"sub_type,omitempty"`

	// 是否已开通
	Subscribe *bool `json:"subscribe,omitempty"`
}

func (o OrderCatalogue) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OrderCatalogue struct{}"
	}

	return strings.Join([]string{"OrderCatalogue", string(data)}, " ")
}
