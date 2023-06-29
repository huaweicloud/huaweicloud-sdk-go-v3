package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddServerNicsResponse Response Object
type AddServerNicsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o AddServerNicsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddServerNicsResponse struct{}"
	}

	return strings.Join([]string{"AddServerNicsResponse", string(data)}, " ")
}
