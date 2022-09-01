package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type BatchUpdateChildNickNamesRequest struct {
	Body *BatchUpdateChildUserNickNamesRequestBody `json:"body,omitempty" xml:"body"`
}

func (o BatchUpdateChildNickNamesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateChildNickNamesRequest struct{}"
	}

	return strings.Join([]string{"BatchUpdateChildNickNamesRequest", string(data)}, " ")
}
