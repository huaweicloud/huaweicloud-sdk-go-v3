package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type QueryDiagnosisResultResultList struct {

	// 内容key。
	Key *string `json:"key,omitempty"`

	// 结果值。
	Value *string `json:"value,omitempty"`

	// 结果值是否需要国际化。
	IsMultiLanguage *bool `json:"is_multi_language,omitempty"`

	// 模块名称。
	ModuleName *string `json:"module_name,omitempty"`

	// 等级。
	Level *string `json:"level,omitempty"`
}

func (o QueryDiagnosisResultResultList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryDiagnosisResultResultList struct{}"
	}

	return strings.Join([]string{"QueryDiagnosisResultResultList", string(data)}, " ")
}
