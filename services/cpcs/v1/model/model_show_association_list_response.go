package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAssociationListResponse Response Object
type ShowAssociationListResponse struct {

	// 满足条件的绑定关系总数
	TotalNum *int32 `json:"total_num,omitempty"`

	// 绑定关系列表
	Result         *[]AssociationInfo `json:"result,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ShowAssociationListResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAssociationListResponse struct{}"
	}

	return strings.Join([]string{"ShowAssociationListResponse", string(data)}, " ")
}
