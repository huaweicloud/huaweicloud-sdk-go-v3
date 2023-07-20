package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateGlobalValueReq 创建全局变量
type CreateGlobalValueReq struct {

	// 变量名称
	VarName string `json:"var_name"`

	// 变量的值
	VarValue string `json:"var_value"`

	// 是否为敏感变量
	IsSensitive *bool `json:"is_sensitive,omitempty"`
}

func (o CreateGlobalValueReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateGlobalValueReq struct{}"
	}

	return strings.Join([]string{"CreateGlobalValueReq", string(data)}, " ")
}
