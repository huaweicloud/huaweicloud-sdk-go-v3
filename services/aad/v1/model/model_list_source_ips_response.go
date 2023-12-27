package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSourceIpsResponse Response Object
type ListSourceIpsResponse struct {

	// 查询高防回源IP段列表
	Ips            *[]string `json:"ips,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListSourceIpsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSourceIpsResponse struct{}"
	}

	return strings.Join([]string{"ListSourceIpsResponse", string(data)}, " ")
}
