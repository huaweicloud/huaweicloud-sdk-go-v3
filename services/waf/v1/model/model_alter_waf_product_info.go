package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AlterWafProductInfo struct {

	// waf规格   - professional：标准版   - enterprise：专业版   ultimate：铂金版
	ResourceSpecCode *string `json:"resource_spec_code,omitempty"`
}

func (o AlterWafProductInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlterWafProductInfo struct{}"
	}

	return strings.Join([]string{"AlterWafProductInfo", string(data)}, " ")
}
