package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListEnterpriseResourceDetail struct {

	// 函数urn
	DetailId *string `json:"detailId,omitempty"`
}

func (o ListEnterpriseResourceDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEnterpriseResourceDetail struct{}"
	}

	return strings.Join([]string{"ListEnterpriseResourceDetail", string(data)}, " ")
}
