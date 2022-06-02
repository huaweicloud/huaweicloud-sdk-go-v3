package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateProjectV4Request struct {

	// devcloud项目的32位id
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
