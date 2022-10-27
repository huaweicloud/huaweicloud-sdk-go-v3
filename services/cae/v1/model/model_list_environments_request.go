package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListEnvironmentsRequest struct {

	// 租户的企业项目id。
	XEnterpriseProjectID *string `json:"X-Enterprise-Project-ID,omitempty"`
}

func (o ListEnvironmentsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEnvironmentsRequest struct{}"
	}

	return strings.Join([]string{"ListEnvironmentsRequest", string(data)}, " ")
}
