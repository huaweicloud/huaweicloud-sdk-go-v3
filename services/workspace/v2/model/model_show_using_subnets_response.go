package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowUsingSubnetsResponse Response Object
type ShowUsingSubnetsResponse struct {

	// 正在被桌面使用的子网ids。
	SubnetIds      *[]string `json:"subnet_ids,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ShowUsingSubnetsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowUsingSubnetsResponse struct{}"
	}

	return strings.Join([]string{"ShowUsingSubnetsResponse", string(data)}, " ")
}
