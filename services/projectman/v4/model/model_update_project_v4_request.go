package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateProjectV4Request struct {
	// devcloud的项目id

	ProjectId string `json:"project_id"`

	Body *UpdateProjectRequestV4 `json:"body,omitempty"`
}

func (o UpdateProjectV4Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateProjectV4Request struct{}"
	}

	return strings.Join([]string{"UpdateProjectV4Request", string(data)}, " ")
}
