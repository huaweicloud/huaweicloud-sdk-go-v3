package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowResourceRequest struct {

	// 资源id.
	ResourceId string `json:"resource_id" xml:"resource_id"`
}

func (o ShowResourceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowResourceRequest struct{}"
	}

	return strings.Join([]string{"ShowResourceRequest", string(data)}, " ")
}
