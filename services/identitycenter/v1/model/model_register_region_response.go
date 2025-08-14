package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RegisterRegionResponse Response Object
type RegisterRegionResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o RegisterRegionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RegisterRegionResponse struct{}"
	}

	return strings.Join([]string{"RegisterRegionResponse", string(data)}, " ")
}
