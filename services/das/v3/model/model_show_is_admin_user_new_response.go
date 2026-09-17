package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowIsAdminUserNewResponse Response Object
type ShowIsAdminUserNewResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ShowIsAdminUserNewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowIsAdminUserNewResponse struct{}"
	}

	return strings.Join([]string{"ShowIsAdminUserNewResponse", string(data)}, " ")
}
