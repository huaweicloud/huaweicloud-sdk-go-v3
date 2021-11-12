package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CompareFaceByUrlRequest struct {
	Body *FaceCompareUrlReq `json:"body,omitempty"`
}

func (o CompareFaceByUrlRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CompareFaceByUrlRequest struct{}"
	}

	return strings.Join([]string{"CompareFaceByUrlRequest", string(data)}, " ")
}
