package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPublicipPoolResponse Response Object
type ListPublicipPoolResponse struct {

	// 功能说明：公网池对象
	PublicipPools *[]PublicipPoolShowResp `json:"publicip_pools,omitempty"`

	// 本次请求的编号
	RequestId *string `json:"request_id,omitempty"`

	PageInfo       *PageInfoOption `json:"page_info,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ListPublicipPoolResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPublicipPoolResponse struct{}"
	}

	return strings.Join([]string{"ListPublicipPoolResponse", string(data)}, " ")
}
