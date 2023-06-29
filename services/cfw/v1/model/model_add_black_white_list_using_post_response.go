package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddBlackWhiteListUsingPostResponse Response Object
type AddBlackWhiteListUsingPostResponse struct {
	Data           *IdObject `json:"data,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o AddBlackWhiteListUsingPostResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddBlackWhiteListUsingPostResponse struct{}"
	}

	return strings.Join([]string{"AddBlackWhiteListUsingPostResponse", string(data)}, " ")
}
