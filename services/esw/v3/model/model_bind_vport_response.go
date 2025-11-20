package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BindVportResponse Response Object
type BindVportResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BindVportResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BindVportResponse struct{}"
	}

	return strings.Join([]string{"BindVportResponse", string(data)}, " ")
}
