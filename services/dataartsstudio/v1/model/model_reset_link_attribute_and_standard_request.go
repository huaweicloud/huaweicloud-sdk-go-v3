package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResetLinkAttributeAndStandardRequest Request Object
type ResetLinkAttributeAndStandardRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	Body *LinkAttributeAndElementVo `json:"body,omitempty"`
}

func (o ResetLinkAttributeAndStandardRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetLinkAttributeAndStandardRequest struct{}"
	}

	return strings.Join([]string{"ResetLinkAttributeAndStandardRequest", string(data)}, " ")
}
