package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeSensitiveInfoResponse Response Object
type ChangeSensitiveInfoResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ChangeSensitiveInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeSensitiveInfoResponse struct{}"
	}

	return strings.Join([]string{"ChangeSensitiveInfoResponse", string(data)}, " ")
}
