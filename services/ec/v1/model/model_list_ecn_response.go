package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEcnResponse Response Object
type ListEcnResponse struct {

	// 企业连接网络列表
	EnterpriseConnectNetworks *[]EcnItem `json:"enterprise_connect_networks,omitempty"`

	PageInfo *PageInfo `json:"page_info,omitempty"`

	// 企业连接网络数量
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListEcnResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEcnResponse struct{}"
	}

	return strings.Join([]string{"ListEcnResponse", string(data)}, " ")
}
