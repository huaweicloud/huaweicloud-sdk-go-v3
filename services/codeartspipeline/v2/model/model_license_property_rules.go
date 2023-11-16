package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type LicensePropertyRules struct {

	// license详情
	Values *[]string `json:"values,omitempty"`

	// 比较规则
	Predicate *string `json:"predicate,omitempty"`
}

func (o LicensePropertyRules) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LicensePropertyRules struct{}"
	}

	return strings.Join([]string{"LicensePropertyRules", string(data)}, " ")
}
