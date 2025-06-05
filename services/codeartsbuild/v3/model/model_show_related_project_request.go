package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRelatedProjectRequest Request Object
type ShowRelatedProjectRequest struct {
}

func (o ShowRelatedProjectRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRelatedProjectRequest struct{}"
	}

	return strings.Join([]string{"ShowRelatedProjectRequest", string(data)}, " ")
}
