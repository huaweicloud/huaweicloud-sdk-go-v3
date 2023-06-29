package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LiveResponse Response Object
type LiveResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o LiveResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LiveResponse struct{}"
	}

	return strings.Join([]string{"LiveResponse", string(data)}, " ")
}
