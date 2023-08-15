package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDashBoardRequest Request Object
type CreateDashBoardRequest struct {
	Body *CreateDashBoardReqBody `json:"body,omitempty"`
}

func (o CreateDashBoardRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDashBoardRequest struct{}"
	}

	return strings.Join([]string{"CreateDashBoardRequest", string(data)}, " ")
}
