package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateWebTamperRaspPathResponse Response Object
type UpdateWebTamperRaspPathResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateWebTamperRaspPathResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateWebTamperRaspPathResponse struct{}"
	}

	return strings.Join([]string{"UpdateWebTamperRaspPathResponse", string(data)}, " ")
}
