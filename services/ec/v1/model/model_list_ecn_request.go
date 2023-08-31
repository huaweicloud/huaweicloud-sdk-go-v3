package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEcnRequest Request Object
type ListEcnRequest struct {

	// 分页查询时每页返回的记录数量
	Limit *int32 `json:"limit,omitempty"`

	// marker标识，请求此marker之后的数据
	Marker *string `json:"marker,omitempty"`

	// 企业项目id
	EnterpriseProjectId *[]string `json:"enterprise_project_id,omitempty"`
}

func (o ListEcnRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEcnRequest struct{}"
	}

	return strings.Join([]string{"ListEcnRequest", string(data)}, " ")
}
