package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateGraphRequest struct {
	Body *CreateGraphReq `json:"body,omitempty" xml:"body"`
}

func (o CreateGraphRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateGraphRequest struct{}"
	}

	return strings.Join([]string{"CreateGraphRequest", string(data)}, " ")
}
