package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckJobInternalRequest Request Object
type CheckJobInternalRequest struct {
}

func (o CheckJobInternalRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckJobInternalRequest struct{}"
	}

	return strings.Join([]string{"CheckJobInternalRequest", string(data)}, " ")
}
