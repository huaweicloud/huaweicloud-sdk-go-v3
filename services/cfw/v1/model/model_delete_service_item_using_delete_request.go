package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteServiceItemUsingDeleteRequest struct {

	// 服务组成员id
	ItemId string `json:"item_id"`
}

func (o DeleteServiceItemUsingDeleteRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteServiceItemUsingDeleteRequest struct{}"
	}

	return strings.Join([]string{"DeleteServiceItemUsingDeleteRequest", string(data)}, " ")
}
