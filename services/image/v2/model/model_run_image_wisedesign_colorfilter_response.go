package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type RunImageWisedesignColorfilterResponse struct {
	Result         *ImageWisedesignColorfilterResponseResult `json:"result,omitempty"`
	HttpStatusCode int                                       `json:"-"`
}

func (o RunImageWisedesignColorfilterResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunImageWisedesignColorfilterResponse struct{}"
	}

	return strings.Join([]string{"RunImageWisedesignColorfilterResponse", string(data)}, " ")
}
