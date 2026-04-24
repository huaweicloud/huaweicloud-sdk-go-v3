package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDcVncResponse Response Object
type UpdateDcVncResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateDcVncResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDcVncResponse struct{}"
	}

	return strings.Join([]string{"UpdateDcVncResponse", string(data)}, " ")
}
