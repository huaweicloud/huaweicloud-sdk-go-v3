package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteBaselineWhiteListResponse Response Object
type DeleteBaselineWhiteListResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteBaselineWhiteListResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteBaselineWhiteListResponse struct{}"
	}

	return strings.Join([]string{"DeleteBaselineWhiteListResponse", string(data)}, " ")
}
