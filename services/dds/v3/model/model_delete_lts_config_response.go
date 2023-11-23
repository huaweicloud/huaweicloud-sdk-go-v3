package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteLtsConfigResponse Response Object
type DeleteLtsConfigResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteLtsConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteLtsConfigResponse struct{}"
	}

	return strings.Join([]string{"DeleteLtsConfigResponse", string(data)}, " ")
}
