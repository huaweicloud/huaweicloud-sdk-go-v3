package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateMigprojectRequest struct {
	Body *PostMigProjectBody `json:"body,omitempty" xml:"body"`
}

func (o CreateMigprojectRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateMigprojectRequest struct{}"
	}

	return strings.Join([]string{"CreateMigprojectRequest", string(data)}, " ")
}
