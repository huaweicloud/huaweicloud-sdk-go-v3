package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OrganizationalPercentageDetail 创建账号、纳管OU、纳管账号状态信息。
type OrganizationalPercentageDetail struct {

	// 进度名称。
	PercentageName *string `json:"percentage_name,omitempty"`

	// 进度完成状态。
	PercentageStatus *string `json:"percentage_status,omitempty"`
}

func (o OrganizationalPercentageDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OrganizationalPercentageDetail struct{}"
	}

	return strings.Join([]string{"OrganizationalPercentageDetail", string(data)}, " ")
}
