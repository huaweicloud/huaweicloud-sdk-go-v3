package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSubnetBandwidthChangeOrderRequestBody 创建云办公带宽变更订单请求体。
type CreateSubnetBandwidthChangeOrderRequestBody struct {

	// 云办公带宽名称。
	BandwidthName *string `json:"bandwidth_name,omitempty"`

	// 变更云办公带宽的带宽大小。
	BandwidthSize string `json:"bandwidth_size"`

	// 企业项目ID。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	ExtendParam *OrderExtendParam `json:"extend_param,omitempty"`
}

func (o CreateSubnetBandwidthChangeOrderRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSubnetBandwidthChangeOrderRequestBody struct{}"
	}

	return strings.Join([]string{"CreateSubnetBandwidthChangeOrderRequestBody", string(data)}, " ")
}
