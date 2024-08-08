package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateImageServerResponse Response Object
type UpdateImageServerResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateImageServerResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateImageServerResponse struct{}"
	}

	return strings.Join([]string{"UpdateImageServerResponse", string(data)}, " ")
}
