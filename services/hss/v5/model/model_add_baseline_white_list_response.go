package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddBaselineWhiteListResponse Response Object
type AddBaselineWhiteListResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o AddBaselineWhiteListResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddBaselineWhiteListResponse struct{}"
	}

	return strings.Join([]string{"AddBaselineWhiteListResponse", string(data)}, " ")
}
