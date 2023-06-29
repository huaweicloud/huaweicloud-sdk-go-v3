package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunImageDescriptionResponse Response Object
type RunImageDescriptionResponse struct {
	Result         *ImageDescriptionResponseResult `json:"result,omitempty"`
	HttpStatusCode int                             `json:"-"`
}

func (o RunImageDescriptionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunImageDescriptionResponse struct{}"
	}

	return strings.Join([]string{"RunImageDescriptionResponse", string(data)}, " ")
}
