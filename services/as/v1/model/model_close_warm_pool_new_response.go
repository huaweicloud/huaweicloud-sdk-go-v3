package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CloseWarmPoolNewResponse Response Object
type CloseWarmPoolNewResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CloseWarmPoolNewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CloseWarmPoolNewResponse struct{}"
	}

	return strings.Join([]string{"CloseWarmPoolNewResponse", string(data)}, " ")
}
