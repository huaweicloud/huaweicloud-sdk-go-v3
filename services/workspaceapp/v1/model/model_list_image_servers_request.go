package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListImageServersRequest Request Object
type ListImageServersRequest struct {

	// 查询的偏移量。
	Offset *int32 `json:"offset,omitempty"`

	// 查询的数量，值区间[1-100]。
	Limit *int32 `json:"limit,omitempty"`

	// 镜像实例名称，支持部分匹配。
	ServerName *string `json:"server_name,omitempty"`

	// 镜像实例唯一标识。
	ServerId *string `json:"server_id,omitempty"`

	// 企业项目ID(字段为空或者0表示使用默认default企业项目)。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o ListImageServersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListImageServersRequest struct{}"
	}

	return strings.Join([]string{"ListImageServersRequest", string(data)}, " ")
}
