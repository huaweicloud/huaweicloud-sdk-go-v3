package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type EnterpriseIdItem struct {

	// 企业项目Id
	EnterpriseId *string `json:"enterprise_id,omitempty"`
}

func (o EnterpriseIdItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnterpriseIdItem struct{}"
	}

	return strings.Join([]string{"EnterpriseIdItem", string(data)}, " ")
}
