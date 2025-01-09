package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeletePacifyWordsResponse Response Object
type BatchDeletePacifyWordsResponse struct {
	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o BatchDeletePacifyWordsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeletePacifyWordsResponse struct{}"
	}

	return strings.Join([]string{"BatchDeletePacifyWordsResponse", string(data)}, " ")
}
