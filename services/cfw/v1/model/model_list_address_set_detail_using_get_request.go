package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListAddressSetDetailUsingGetRequest struct {

	// 地址组id
	SetId string `json:"set_id"`
}

func (o ListAddressSetDetailUsingGetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAddressSetDetailUsingGetRequest struct{}"
	}

	return strings.Join([]string{"ListAddressSetDetailUsingGetRequest", string(data)}, " ")
}
