package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSubjectNewRequest Request Object
type CreateSubjectNewRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	Body *SubjectParamsVo `json:"body,omitempty"`
}

func (o CreateSubjectNewRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSubjectNewRequest struct{}"
	}

	return strings.Join([]string{"CreateSubjectNewRequest", string(data)}, " ")
}
