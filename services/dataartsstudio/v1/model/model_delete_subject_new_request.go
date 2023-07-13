package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteSubjectNewRequest Request Object
type DeleteSubjectNewRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	Body *IdsParam `json:"body,omitempty"`
}

func (o DeleteSubjectNewRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSubjectNewRequest struct{}"
	}

	return strings.Join([]string{"DeleteSubjectNewRequest", string(data)}, " ")
}
