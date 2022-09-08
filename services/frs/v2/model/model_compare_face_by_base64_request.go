package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CompareFaceByBase64Request struct {
	Body *FaceCompareBase64Req `json:"body,omitempty"`
}

func (o CompareFaceByBase64Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CompareFaceByBase64Request struct{}"
	}

	return strings.Join([]string{"CompareFaceByBase64Request", string(data)}, " ")
}
