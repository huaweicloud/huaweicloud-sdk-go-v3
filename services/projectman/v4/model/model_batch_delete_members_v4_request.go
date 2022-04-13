package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type BatchDeleteMembersV4Request struct {
	// devcloud的项目id

	ProjectId string `json:"project_id"`

	Body *BatchDeleteMembersV4RequestBody `json:"body,omitempty"`
}

func (o BatchDeleteMembersV4Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteMembersV4Request struct{}"
	}

	return strings.Join([]string{"BatchDeleteMembersV4Request", string(data)}, " ")
}
