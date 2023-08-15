package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateGraphRequest Request Object
type CreateGraphRequest struct {
	Body *CreateGraphReq `json:"body,omitempty"`
}

func (o CreateGraphRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateGraphRequest struct{}"
	}

	return strings.Join([]string{"CreateGraphRequest", string(data)}, " ")
}
