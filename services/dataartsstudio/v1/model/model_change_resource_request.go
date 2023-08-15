package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeResourceRequest Request Object
type ChangeResourceRequest struct {
	Body *ApigChangeResourceReq `json:"body,omitempty"`
}

func (o ChangeResourceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeResourceRequest struct{}"
	}

	return strings.Join([]string{"ChangeResourceRequest", string(data)}, " ")
}
