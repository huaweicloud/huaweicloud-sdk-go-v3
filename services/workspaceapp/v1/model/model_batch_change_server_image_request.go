package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchChangeServerImageRequest Request Object
type BatchChangeServerImageRequest struct {
	Body *BatchChangeServerImageReq `json:"body,omitempty"`
}

func (o BatchChangeServerImageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchChangeServerImageRequest struct{}"
	}

	return strings.Join([]string{"BatchChangeServerImageRequest", string(data)}, " ")
}
