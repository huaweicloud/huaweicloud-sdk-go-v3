package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListVpcInfoRequest Request Object
type ListVpcInfoRequest struct {
}

func (o ListVpcInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVpcInfoRequest struct{}"
	}

	return strings.Join([]string{"ListVpcInfoRequest", string(data)}, " ")
}
