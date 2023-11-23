package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SaveLtsConfigsResponse Response Object
type SaveLtsConfigsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o SaveLtsConfigsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SaveLtsConfigsResponse struct{}"
	}

	return strings.Join([]string{"SaveLtsConfigsResponse", string(data)}, " ")
}
