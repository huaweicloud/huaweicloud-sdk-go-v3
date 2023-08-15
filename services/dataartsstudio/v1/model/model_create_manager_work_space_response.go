package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateManagerWorkSpaceResponse Response Object
type CreateManagerWorkSpaceResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateManagerWorkSpaceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateManagerWorkSpaceResponse struct{}"
	}

	return strings.Join([]string{"CreateManagerWorkSpaceResponse", string(data)}, " ")
}
