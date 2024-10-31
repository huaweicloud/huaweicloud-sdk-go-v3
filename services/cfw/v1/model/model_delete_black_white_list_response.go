package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteBlackWhiteListResponse Response Object
type DeleteBlackWhiteListResponse struct {
	Data           *BlackWhiteListId `json:"data,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o DeleteBlackWhiteListResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteBlackWhiteListResponse struct{}"
	}

	return strings.Join([]string{"DeleteBlackWhiteListResponse", string(data)}, " ")
}
