package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateMembesRoleV4Request Request Object
type UpdateMembesRoleV4Request struct {

	// devcloud项目的32位id
	ProjectId string `json:"project_id"`

	Body *UpdateMembesRoleV4RequestBody `json:"body,omitempty"`
}

func (o UpdateMembesRoleV4Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateMembesRoleV4Request struct{}"
	}

	return strings.Join([]string{"UpdateMembesRoleV4Request", string(data)}, " ")
}
