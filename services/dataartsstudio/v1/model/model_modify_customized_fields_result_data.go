package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModifyCustomizedFieldsResultData data，统一的返回结果的最外层数据结构。
type ModifyCustomizedFieldsResultData struct {

	// 数据连接信息数组
	Value *[]CustomizedFieldsVo `json:"value,omitempty"`
}

func (o ModifyCustomizedFieldsResultData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyCustomizedFieldsResultData struct{}"
	}

	return strings.Join([]string{"ModifyCustomizedFieldsResultData", string(data)}, " ")
}
