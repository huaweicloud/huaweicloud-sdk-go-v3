package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteStandardRequest Request Object
type DeleteStandardRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	Body *IdsParam `json:"body,omitempty"`
}

func (o DeleteStandardRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteStandardRequest struct{}"
	}

	return strings.Join([]string{"DeleteStandardRequest", string(data)}, " ")
}
