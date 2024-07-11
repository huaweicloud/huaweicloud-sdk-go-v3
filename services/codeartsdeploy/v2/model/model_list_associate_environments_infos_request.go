package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAssociateEnvironmentsInfosRequest Request Object
type ListAssociateEnvironmentsInfosRequest struct {

	// 主机集群id
	GroupId string `json:"group_id"`

	// 页码
	PageIndex *int32 `json:"page_index,omitempty"`

	// 每页查询条数
	PageSize *int32 `json:"page_size,omitempty"`
}

func (o ListAssociateEnvironmentsInfosRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAssociateEnvironmentsInfosRequest struct{}"
	}

	return strings.Join([]string{"ListAssociateEnvironmentsInfosRequest", string(data)}, " ")
}
