package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunTtsRequest Request Object
type RunTtsRequest struct {
	Body *PostCustomTtsReq `json:"body,omitempty"`
}

func (o RunTtsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunTtsRequest struct{}"
	}

	return strings.Join([]string{"RunTtsRequest", string(data)}, " ")
}
