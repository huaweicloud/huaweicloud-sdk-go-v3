package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteAddressSetInfoUsingDeleteRequest struct {

	// 地址组id
	SetId string `json:"set_id"`
}

func (o DeleteAddressSetInfoUsingDeleteRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAddressSetInfoUsingDeleteRequest struct{}"
	}

	return strings.Join([]string{"DeleteAddressSetInfoUsingDeleteRequest", string(data)}, " ")
}
