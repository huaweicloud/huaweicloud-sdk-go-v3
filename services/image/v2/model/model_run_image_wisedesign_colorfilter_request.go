package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type RunImageWisedesignColorfilterRequest struct {
	Body *ImageWisedesignColorfilterReq `json:"body,omitempty"`
}

func (o RunImageWisedesignColorfilterRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunImageWisedesignColorfilterRequest struct{}"
	}

	return strings.Join([]string{"RunImageWisedesignColorfilterRequest", string(data)}, " ")
}
