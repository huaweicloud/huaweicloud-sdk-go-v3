package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeBaselineWhiteListResponse Response Object
type ChangeBaselineWhiteListResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ChangeBaselineWhiteListResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeBaselineWhiteListResponse struct{}"
	}

	return strings.Join([]string{"ChangeBaselineWhiteListResponse", string(data)}, " ")
}
