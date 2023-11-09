package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListVgwsRequest Request Object
type ListVgwsRequest struct {

	// vgw ID
	VgwId *string `json:"vgw_id,omitempty"`

	// 企业项目id
	EnterpriseProjectId *[]string `json:"enterprise_project_id,omitempty"`
}

func (o ListVgwsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVgwsRequest struct{}"
	}

	return strings.Join([]string{"ListVgwsRequest", string(data)}, " ")
}
