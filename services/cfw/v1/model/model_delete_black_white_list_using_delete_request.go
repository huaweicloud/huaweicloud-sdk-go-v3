package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteBlackWhiteListUsingDeleteRequest struct {

	// 黑白名单id
	ListId string `json:"list_id"`
}

func (o DeleteBlackWhiteListUsingDeleteRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteBlackWhiteListUsingDeleteRequest struct{}"
	}

	return strings.Join([]string{"DeleteBlackWhiteListUsingDeleteRequest", string(data)}, " ")
}
