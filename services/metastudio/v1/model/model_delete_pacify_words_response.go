package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeletePacifyWordsResponse Response Object
type DeletePacifyWordsResponse struct {
	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeletePacifyWordsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePacifyWordsResponse struct{}"
	}

	return strings.Join([]string{"DeletePacifyWordsResponse", string(data)}, " ")
}
