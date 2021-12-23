package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListBlockchainsRequest struct {
}

func (o ListBlockchainsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBlockchainsRequest struct{}"
	}

	return strings.Join([]string{"ListBlockchainsRequest", string(data)}, " ")
}
