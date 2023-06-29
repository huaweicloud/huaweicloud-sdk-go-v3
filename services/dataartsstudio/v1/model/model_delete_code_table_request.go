package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteCodeTableRequest Request Object
type DeleteCodeTableRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	Body *IdsParam `json:"body,omitempty"`
}

func (o DeleteCodeTableRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteCodeTableRequest struct{}"
	}

	return strings.Join([]string{"DeleteCodeTableRequest", string(data)}, " ")
}
