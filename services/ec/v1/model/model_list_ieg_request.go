package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListIegRequest Request Object
type ListIegRequest struct {

	// 分页查询时每页返回的记录数量
	Limit *int32 `json:"limit,omitempty"`

	// marker标识，请求此marker之后的数据
	Marker *string `json:"marker,omitempty"`

	// 企业项目id
	EnterpriseProjectId *[]string `json:"enterprise_project_id,omitempty"`
}

func (o ListIegRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListIegRequest struct{}"
	}

	return strings.Join([]string{"ListIegRequest", string(data)}, " ")
}
