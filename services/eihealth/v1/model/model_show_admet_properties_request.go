package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAdmetPropertiesRequest Request Object
type ShowAdmetPropertiesRequest struct {
	Body *AdmetRequest `json:"body,omitempty"`
}

func (o ShowAdmetPropertiesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAdmetPropertiesRequest struct{}"
	}

	return strings.Join([]string{"ShowAdmetPropertiesRequest", string(data)}, " ")
}
