package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPointsResponse Response Object
type ShowPointsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ShowPointsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPointsResponse struct{}"
	}

	return strings.Join([]string{"ShowPointsResponse", string(data)}, " ")
}
