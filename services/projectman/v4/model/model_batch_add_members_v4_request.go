package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type BatchAddMembersV4Request struct {

	// devcloud项目的32位id
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
