package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSessionTypesRequest Request Object
type ShowSessionTypesRequest struct {
	Body *QuerySessionTypesReq `json:"body,omitempty"`
}

func (o ShowSessionTypesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSessionTypesRequest struct{}"
	}

	return strings.Join([]string{"ShowSessionTypesRequest", string(data)}, " ")
}
