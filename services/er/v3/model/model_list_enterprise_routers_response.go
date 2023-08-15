package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEnterpriseRoutersResponse Response Object
type ListEnterpriseRoutersResponse struct {

	// 企业路由器列表
	Instances *[]EnterpriseRouter `json:"instances,omitempty"`

	PageInfo *PageInfo `json:"page_info,omitempty"`

	// 请求ID
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListEnterpriseRoutersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEnterpriseRoutersResponse struct{}"
	}

	return strings.Join([]string{"ListEnterpriseRoutersResponse", string(data)}, " ")
}
