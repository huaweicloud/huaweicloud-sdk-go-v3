package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListOrganizationTreeResponse Response Object
type ListOrganizationTreeResponse struct {

	// **参数解释**: 总数 **取值范围**: 最小值0，最大值2147483647
	TotalNum *int32 `json:"total_num,omitempty"`

	// **参数解释** 存储查询到的组织树节点详细信息列表； **取值范围** 数组长度0-当前查询的默认limit值（常规默认10条，未指定时按平台默认规则），数组元素为OrganizationNodeResponseInfo对象
	DataList *[]OrganizationNodeResponseInfo `json:"data_list,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListOrganizationTreeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListOrganizationTreeResponse struct{}"
	}

	return strings.Join([]string{"ListOrganizationTreeResponse", string(data)}, " ")
}
