package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAssetResourceRequest Request Object
type CreateAssetResourceRequest struct {
	Body *CreateAssetResourceReq `json:"body,omitempty"`
}

func (o CreateAssetResourceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAssetResourceRequest struct{}"
	}

	return strings.Join([]string{"CreateAssetResourceRequest", string(data)}, " ")
}
