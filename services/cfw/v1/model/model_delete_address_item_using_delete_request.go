package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteAddressItemUsingDeleteRequest struct {

	// 地址组成员id
	ItemId string `json:"item_id"`
}

func (o DeleteAddressItemUsingDeleteRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAddressItemUsingDeleteRequest struct{}"
	}

	return strings.Join([]string{"DeleteAddressItemUsingDeleteRequest", string(data)}, " ")
}
