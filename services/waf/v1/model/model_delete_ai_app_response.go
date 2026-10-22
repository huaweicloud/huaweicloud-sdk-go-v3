package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAiAppResponse Response Object
type DeleteAiAppResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteAiAppResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAiAppResponse struct{}"
	}

	return strings.Join([]string{"DeleteAiAppResponse", string(data)}, " ")
}
