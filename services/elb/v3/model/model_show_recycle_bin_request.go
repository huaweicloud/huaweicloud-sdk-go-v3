package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRecycleBinRequest Request Object
type ShowRecycleBinRequest struct {
}

func (o ShowRecycleBinRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRecycleBinRequest struct{}"
	}

	return strings.Join([]string{"ShowRecycleBinRequest", string(data)}, " ")
}
