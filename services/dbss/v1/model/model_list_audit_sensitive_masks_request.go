package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAuditSensitiveMasksRequest Request Object
type ListAuditSensitiveMasksRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 偏移量
	Offset *string `json:"offset,omitempty"`

	// 查询记录数
	Limit *string `json:"limit,omitempty"`
}

func (o ListAuditSensitiveMasksRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAuditSensitiveMasksRequest struct{}"
	}

	return strings.Join([]string{"ListAuditSensitiveMasksRequest", string(data)}, " ")
}
