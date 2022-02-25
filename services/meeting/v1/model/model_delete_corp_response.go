package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteCorpResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteCorpResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteCorpResponse struct{}"
	}

	return strings.Join([]string{"DeleteCorpResponse", string(data)}, " ")
}
