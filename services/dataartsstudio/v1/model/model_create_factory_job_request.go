package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateFactoryJobRequest Request Object
type CreateFactoryJobRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	Body *JobInfoRequest `json:"body,omitempty"`
}

func (o CreateFactoryJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateFactoryJobRequest struct{}"
	}

	return strings.Join([]string{"CreateFactoryJobRequest", string(data)}, " ")
}
