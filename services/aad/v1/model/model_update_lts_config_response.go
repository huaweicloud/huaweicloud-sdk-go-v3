package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateLtsConfigResponse Response Object
type UpdateLtsConfigResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateLtsConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateLtsConfigResponse struct{}"
	}

	return strings.Join([]string{"UpdateLtsConfigResponse", string(data)}, " ")
}
