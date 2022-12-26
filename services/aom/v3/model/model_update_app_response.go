package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateAppResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateAppResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAppResponse struct{}"
	}

	return strings.Join([]string{"UpdateAppResponse", string(data)}, " ")
}
