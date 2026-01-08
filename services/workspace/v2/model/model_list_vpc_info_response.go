package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListVpcInfoResponse Response Object
type ListVpcInfoResponse struct {

	// vpc信息列表。
	Vpcs           *[]VpcInfo `json:"vpcs,omitempty"`
	HttpStatusCode int        `json:"-"`
}

func (o ListVpcInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVpcInfoResponse struct{}"
	}

	return strings.Join([]string{"ListVpcInfoResponse", string(data)}, " ")
}
