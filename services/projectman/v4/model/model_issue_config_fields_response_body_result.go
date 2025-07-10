package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IssueConfigFieldsResponseBodyResult 查询结果
type IssueConfigFieldsResponseBodyResult struct {

	// 字段配置
	ConfigFields *[]IssueConfigFieldsResponseBodyResultConfigFields `json:"config_fields,omitempty"`
}

func (o IssueConfigFieldsResponseBodyResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IssueConfigFieldsResponseBodyResult struct{}"
	}

	return strings.Join([]string{"IssueConfigFieldsResponseBodyResult", string(data)}, " ")
}
