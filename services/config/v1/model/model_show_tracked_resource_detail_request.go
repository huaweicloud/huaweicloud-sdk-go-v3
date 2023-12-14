package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTrackedResourceDetailRequest Request Object
type ShowTrackedResourceDetailRequest struct {

	// 资源ID
	ResourceId string `json:"resource_id"`
}

func (o ShowTrackedResourceDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTrackedResourceDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowTrackedResourceDetailRequest", string(data)}, " ")
}
