package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateBlackWhiteListUsingPutRequest struct {

	// 黑白名单列表id
	ListId string `json:"list_id"`

	Body *UpdateBlackWhiteListDto `json:"body,omitempty"`
}

func (o UpdateBlackWhiteListUsingPutRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateBlackWhiteListUsingPutRequest struct{}"
	}

	return strings.Join([]string{"UpdateBlackWhiteListUsingPutRequest", string(data)}, " ")
}
