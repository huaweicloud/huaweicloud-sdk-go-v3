package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AddProtectRequest struct {
	AccessLevel *AddProtectAccessLevel `json:"access_level"`
}

func (o AddProtectRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddProtectRequest struct{}"
	}

	return strings.Join([]string{"AddProtectRequest", string(data)}, " ")
}
