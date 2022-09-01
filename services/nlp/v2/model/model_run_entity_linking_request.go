package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type RunEntityLinkingRequest struct {
	Body *PostEntityLinkingRequest `json:"body,omitempty" xml:"body"`
}

func (o RunEntityLinkingRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunEntityLinkingRequest struct{}"
	}

	return strings.Join([]string{"RunEntityLinkingRequest", string(data)}, " ")
}
