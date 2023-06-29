package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchUpdateChildNickNamesRequest Request Object
type BatchUpdateChildNickNamesRequest struct {
	Body *BatchUpdateChildUserNickNamesRequestBody `json:"body,omitempty"`
}

func (o BatchUpdateChildNickNamesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateChildNickNamesRequest struct{}"
	}

	return strings.Join([]string{"BatchUpdateChildNickNamesRequest", string(data)}, " ")
}
