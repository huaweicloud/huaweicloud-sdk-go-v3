package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteLtsConfigsResponse Response Object
type DeleteLtsConfigsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteLtsConfigsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteLtsConfigsResponse struct{}"
	}

	return strings.Join([]string{"DeleteLtsConfigsResponse", string(data)}, " ")
}
