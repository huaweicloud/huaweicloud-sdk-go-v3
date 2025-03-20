package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GetServiceLinkedAgencyDeletionStatusV5Response Response Object
type GetServiceLinkedAgencyDeletionStatusV5Response struct {
	Status *DeletionTaskStatus `json:"status,omitempty"`

	// 删除失败的原因。
	Reason *string `json:"reason,omitempty"`

	// 该服务关联委托正在被使用的场景列表。
	AgencyUsageList *[]AgencyUsage `json:"agency_usage_list,omitempty"`
	HttpStatusCode  int            `json:"-"`
}

func (o GetServiceLinkedAgencyDeletionStatusV5Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetServiceLinkedAgencyDeletionStatusV5Response struct{}"
	}

	return strings.Join([]string{"GetServiceLinkedAgencyDeletionStatusV5Response", string(data)}, " ")
}
