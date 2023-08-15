package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowResourceDetailRequest Request Object
type ShowResourceDetailRequest struct {

	// 资源ID
	ResourceId string `json:"resource_id"`
}

func (o ShowResourceDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowResourceDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowResourceDetailRequest", string(data)}, " ")
}
