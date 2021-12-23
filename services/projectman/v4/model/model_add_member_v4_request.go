package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type AddMemberV4Request struct {
	// devcloud的项目id

	ProjectId string `json:"project_id"`

	Body *AddMemberRequestV4 `json:"body,omitempty"`
}

func (o AddMemberV4Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddMemberV4Request struct{}"
	}

	return strings.Join([]string{"AddMemberV4Request", string(data)}, " ")
}
