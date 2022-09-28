package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type RebootNodeResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o RebootNodeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RebootNodeResponse struct{}"
	}

	return strings.Join([]string{"RebootNodeResponse", string(data)}, " ")
}
