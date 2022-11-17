package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type AddBlackWhiteListUsingPostRequest struct {
	Body *AddBlackWhiteListDto `json:"body,omitempty"`
}

func (o AddBlackWhiteListUsingPostRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddBlackWhiteListUsingPostRequest struct{}"
	}

	return strings.Join([]string{"AddBlackWhiteListUsingPostRequest", string(data)}, " ")
}
