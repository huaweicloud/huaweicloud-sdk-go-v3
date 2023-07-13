package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateSubjectNewRequest Request Object
type UpdateSubjectNewRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	Body *SubjectParamsVo `json:"body,omitempty"`
}

func (o UpdateSubjectNewRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSubjectNewRequest struct{}"
	}

	return strings.Join([]string{"UpdateSubjectNewRequest", string(data)}, " ")
}
