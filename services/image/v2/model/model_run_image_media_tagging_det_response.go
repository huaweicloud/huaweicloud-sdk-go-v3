package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunImageMediaTaggingDetResponse Response Object
type RunImageMediaTaggingDetResponse struct {
	Result         *ImageMediaTaggingDetResponseResult `json:"result,omitempty"`
	HttpStatusCode int                                 `json:"-"`
}

func (o RunImageMediaTaggingDetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunImageMediaTaggingDetResponse struct{}"
	}

	return strings.Join([]string{"RunImageMediaTaggingDetResponse", string(data)}, " ")
}
