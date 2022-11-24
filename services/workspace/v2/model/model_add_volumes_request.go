package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type AddVolumesRequest struct {
	Body *AddDesktopsVolumesReq `json:"body,omitempty"`
}

func (o AddVolumesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddVolumesRequest struct{}"
	}

	return strings.Join([]string{"AddVolumesRequest", string(data)}, " ")
}
