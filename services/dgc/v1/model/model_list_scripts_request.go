package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListScriptsRequest struct {
}

func (o ListScriptsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListScriptsRequest struct{}"
	}

	return strings.Join([]string{"ListScriptsRequest", string(data)}, " ")
}
