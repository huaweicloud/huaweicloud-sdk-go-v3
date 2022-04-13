package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
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
