package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateSpeedResponse Response Object
type UpdateSpeedResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateSpeedResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSpeedResponse struct{}"
	}

	return strings.Join([]string{"UpdateSpeedResponse", string(data)}, " ")
}
