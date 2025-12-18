package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateLongAkskConfigResponse Response Object
type UpdateLongAkskConfigResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateLongAkskConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateLongAkskConfigResponse struct{}"
	}

	return strings.Join([]string{"UpdateLongAkskConfigResponse", string(data)}, " ")
}
