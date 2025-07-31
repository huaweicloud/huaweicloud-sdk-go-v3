package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateStreamForbiddenOnceResponse Response Object
type CreateStreamForbiddenOnceResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateStreamForbiddenOnceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateStreamForbiddenOnceResponse struct{}"
	}

	return strings.Join([]string{"CreateStreamForbiddenOnceResponse", string(data)}, " ")
}
