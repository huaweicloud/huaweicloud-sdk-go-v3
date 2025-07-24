package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowInstanceStatusRequest Request Object
type ShowInstanceStatusRequest struct {

	// imetal instance id
	Id string `json:"id"`
}

func (o ShowInstanceStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceStatusRequest struct{}"
	}

	return strings.Join([]string{"ShowInstanceStatusRequest", string(data)}, " ")
}
