package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListErrorServersRequest Request Object
type ListErrorServersRequest struct {

	// 每一页记录的错误数量
	Limit *int32 `json:"limit,omitempty"`

	// 偏移量
	Offset int32 `json:"offset"`

	// 需要查询的企业项目ID
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o ListErrorServersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListErrorServersRequest struct{}"
	}

	return strings.Join([]string{"ListErrorServersRequest", string(data)}, " ")
}
