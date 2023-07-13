package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteSubjectRequest Request Object
type DeleteSubjectRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	Body *IdsParam `json:"body,omitempty"`
}

func (o DeleteSubjectRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSubjectRequest struct{}"
	}

	return strings.Join([]string{"DeleteSubjectRequest", string(data)}, " ")
}
