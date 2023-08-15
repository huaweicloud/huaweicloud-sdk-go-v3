package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunEntityLinkingRequest Request Object
type RunEntityLinkingRequest struct {
	Body *PostEntityLinkingRequest `json:"body,omitempty"`
}

func (o RunEntityLinkingRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunEntityLinkingRequest struct{}"
	}

	return strings.Join([]string{"RunEntityLinkingRequest", string(data)}, " ")
}
