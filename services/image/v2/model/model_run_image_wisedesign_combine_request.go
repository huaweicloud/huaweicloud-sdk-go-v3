package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type RunImageWisedesignCombineRequest struct {
	Body *ImageWisedesignCombineReq `json:"body,omitempty"`
}

func (o RunImageWisedesignCombineRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunImageWisedesignCombineRequest struct{}"
	}

	return strings.Join([]string{"RunImageWisedesignCombineRequest", string(data)}, " ")
}
