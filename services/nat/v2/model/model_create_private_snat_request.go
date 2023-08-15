package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePrivateSnatRequest Request Object
type CreatePrivateSnatRequest struct {
	Body *CreatePrivateSnatOptionBody `json:"body,omitempty"`
}

func (o CreatePrivateSnatRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePrivateSnatRequest struct{}"
	}

	return strings.Join([]string{"CreatePrivateSnatRequest", string(data)}, " ")
}
