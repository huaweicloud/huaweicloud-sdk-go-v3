package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateStreamForbiddenOnceRequest Request Object
type CreateStreamForbiddenOnceRequest struct {
	Body *StreamForbiddenOnceSetting `json:"body,omitempty"`
}

func (o CreateStreamForbiddenOnceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateStreamForbiddenOnceRequest struct{}"
	}

	return strings.Join([]string{"CreateStreamForbiddenOnceRequest", string(data)}, " ")
}
