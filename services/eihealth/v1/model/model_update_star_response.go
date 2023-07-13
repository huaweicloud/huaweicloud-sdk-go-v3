package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateStarResponse Response Object
type UpdateStarResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateStarResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateStarResponse struct{}"
	}

	return strings.Join([]string{"UpdateStarResponse", string(data)}, " ")
}
