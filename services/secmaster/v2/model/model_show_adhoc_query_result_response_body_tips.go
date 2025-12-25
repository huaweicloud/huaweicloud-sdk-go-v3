package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAdhocQueryResultResponseBodyTips 相关信息
type ShowAdhocQueryResultResponseBodyTips struct {

	// 错误码
	ErrorCode *string `json:"error_code,omitempty"`

	// 错误键
	ErrorKey *string `json:"error_key,omitempty"`

	// 模块名称
	ModuleName *string `json:"module_name,omitempty"`
}

func (o ShowAdhocQueryResultResponseBodyTips) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAdhocQueryResultResponseBodyTips struct{}"
	}

	return strings.Join([]string{"ShowAdhocQueryResultResponseBodyTips", string(data)}, " ")
}
