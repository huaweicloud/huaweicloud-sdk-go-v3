package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ExpandVolumesRequest struct {
	Body *ExpandDesktopsVolumesReq `json:"body,omitempty"`
}

func (o ExpandVolumesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExpandVolumesRequest struct{}"
	}

	return strings.Join([]string{"ExpandVolumesRequest", string(data)}, " ")
}
