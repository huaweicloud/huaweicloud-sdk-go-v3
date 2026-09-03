package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateSearchPathFlagResponse Response Object
type UpdateSearchPathFlagResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateSearchPathFlagResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSearchPathFlagResponse struct{}"
	}

	return strings.Join([]string{"UpdateSearchPathFlagResponse", string(data)}, " ")
}
