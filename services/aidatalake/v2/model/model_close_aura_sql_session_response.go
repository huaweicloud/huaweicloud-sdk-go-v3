package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CloseAuraSqlSessionResponse Response Object
type CloseAuraSqlSessionResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CloseAuraSqlSessionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CloseAuraSqlSessionResponse struct{}"
	}

	return strings.Join([]string{"CloseAuraSqlSessionResponse", string(data)}, " ")
}
