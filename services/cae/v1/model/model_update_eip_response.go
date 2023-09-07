package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateEipResponse Response Object
type UpdateEipResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateEipResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEipResponse struct{}"
	}

	return strings.Join([]string{"UpdateEipResponse", string(data)}, " ")
}
