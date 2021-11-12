package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SeverityV2Do struct {
	// 语言

	Language *string `json:"language,omitempty"`
	// 严重性名称

	SeverityName *string `json:"severity_name,omitempty"`
	// 严重性id

	SeverityId *string `json:"severity_id,omitempty"`
}

func (o SeverityV2Do) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SeverityV2Do struct{}"
	}

	return strings.Join([]string{"SeverityV2Do", string(data)}, " ")
}
