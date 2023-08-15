package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatefavoriteRequest Request Object
type CreatefavoriteRequest struct {
	Body *CreatefavoriteReqbody `json:"body,omitempty"`
}

func (o CreatefavoriteRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatefavoriteRequest struct{}"
	}

	return strings.Join([]string{"CreatefavoriteRequest", string(data)}, " ")
}
