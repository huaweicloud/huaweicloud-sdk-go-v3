package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDrugJobReq 更新药物作业请求体
type UpdateDrugJobReq struct {

	// 标签，取值范围[0,5]，单个标签最大长度32字符，支持中文、字母、数字、空格、下划线和中划线，且不能以空格开头或者尾
	Labels *[]string `json:"labels,omitempty"`
}

func (o UpdateDrugJobReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDrugJobReq struct{}"
	}

	return strings.Join([]string{"UpdateDrugJobReq", string(data)}, " ")
}
