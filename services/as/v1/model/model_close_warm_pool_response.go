package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CloseWarmPoolResponse Response Object
type CloseWarmPoolResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CloseWarmPoolResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CloseWarmPoolResponse struct{}"
	}

	return strings.Join([]string{"CloseWarmPoolResponse", string(data)}, " ")
}
