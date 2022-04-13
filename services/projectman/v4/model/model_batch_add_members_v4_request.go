package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type BatchAddMembersV4Request struct {
	// devcloud的项目id

	ProjectId string `json:"project_id"`

	Body *BatchAddMembersV4RequestBody `json:"body,omitempty"`
}

func (o BatchAddMembersV4Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchAddMembersV4Request struct{}"
	}

	return strings.Join([]string{"BatchAddMembersV4Request", string(data)}, " ")
}
