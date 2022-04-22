package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateInstanceRemarkRequestBody struct {

	// 实例备注内容。 长度不能超过64位，不支持回车和>!<\"&'=特殊字符。
	Remark string `json:"remark"`
}

func (o UpdateInstanceRemarkRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInstanceRemarkRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateInstanceRemarkRequestBody", string(data)}, " ")
}
