package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAppResponse Response Object
type UpdateAppResponse struct {
	App            *AppResp `json:"app,omitempty"`
	HttpStatusCode int      `json:"-"`
}

func (o UpdateAppResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAppResponse struct{}"
	}

	return strings.Join([]string{"UpdateAppResponse", string(data)}, " ")
}
