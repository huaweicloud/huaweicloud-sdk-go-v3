package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateGroupV5Request Request Object
type CreateGroupV5Request struct {
	Body *CreateGroupReqBody `json:"body,omitempty"`
}

func (o CreateGroupV5Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateGroupV5Request struct{}"
	}

	return strings.Join([]string{"CreateGroupV5Request", string(data)}, " ")
}
