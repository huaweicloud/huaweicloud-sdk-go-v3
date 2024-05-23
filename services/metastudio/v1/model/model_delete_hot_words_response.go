package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteHotWordsResponse Response Object
type DeleteHotWordsResponse struct {
	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteHotWordsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteHotWordsResponse struct{}"
	}

	return strings.Join([]string{"DeleteHotWordsResponse", string(data)}, " ")
}
