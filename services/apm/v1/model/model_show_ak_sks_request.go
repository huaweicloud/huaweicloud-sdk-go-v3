package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowAkSksRequest struct {
}

func (o ShowAkSksRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAkSksRequest struct{}"
	}

	return strings.Join([]string{"ShowAkSksRequest", string(data)}, " ")
}
