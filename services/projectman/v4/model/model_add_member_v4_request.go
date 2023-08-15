package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddMemberV4Request Request Object
type AddMemberV4Request struct {

	// devcloud项目的32位id
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
