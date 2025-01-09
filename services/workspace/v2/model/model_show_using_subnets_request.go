package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowUsingSubnetsRequest Request Object
type ShowUsingSubnetsRequest struct {

	// 子网id列表逗号分隔。
	SubnetIds *string `json:"subnet_ids,omitempty"`
}

func (o ShowUsingSubnetsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowUsingSubnetsRequest struct{}"
	}

	return strings.Join([]string{"ShowUsingSubnetsRequest", string(data)}, " ")
}
