package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateHostProtectStatusResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateHostProtectStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateHostProtectStatusResponse struct{}"
	}

	return strings.Join([]string{"UpdateHostProtectStatusResponse", string(data)}, " ")
}
