package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListComponentsRequest struct {

	// 环境id。
	XEnvironmentID string `json:"X-Environment-ID"`

	// 租户的企业项目id。
	XEnterpriseProjectID *string `json:"X-Enterprise-Project-ID,omitempty"`

	// 限制本次返回结果的条数。
	Limit *string `json:"limit,omitempty"`

	// 分页偏移位，查询起始位置。
	Offset *string `json:"offset,omitempty"`

	// 应用id。
	ApplicationId string `json:"application_id"`
}

func (o ListComponentsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListComponentsRequest struct{}"
	}

	return strings.Join([]string{"ListComponentsRequest", string(data)}, " ")
}
