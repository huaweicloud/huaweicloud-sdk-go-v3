package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateStandardRequest Request Object
type CreateStandardRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	Body *StandElementValueVoList `json:"body,omitempty"`
}

func (o CreateStandardRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateStandardRequest struct{}"
	}

	return strings.Join([]string{"CreateStandardRequest", string(data)}, " ")
}
