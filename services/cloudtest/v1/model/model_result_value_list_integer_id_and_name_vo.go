package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResultValueListIntegerIdAndNameVo 请求的返回的数据对象
type ResultValueListIntegerIdAndNameVo struct {
	Value *[]IntegerIdAndNameVo `json:"value,omitempty"`
}

func (o ResultValueListIntegerIdAndNameVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResultValueListIntegerIdAndNameVo struct{}"
	}

	return strings.Join([]string{"ResultValueListIntegerIdAndNameVo", string(data)}, " ")
}
