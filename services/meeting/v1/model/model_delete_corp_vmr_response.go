package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteCorpVmrResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteCorpVmrResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteCorpVmrResponse struct{}"
	}

	return strings.Join([]string{"DeleteCorpVmrResponse", string(data)}, " ")
}
