package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type HbaseModifySettingV2 struct {

	// 待修改的参数名
	ParmName string `json:"parm_name"`

	// 设置的参数值
	NewValue string `json:"new_value"`

	// 参数ID，可不传
	Id *string `json:"id,omitempty"`
}

func (o HbaseModifySettingV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HbaseModifySettingV2 struct{}"
	}

	return strings.Join([]string{"HbaseModifySettingV2", string(data)}, " ")
}
