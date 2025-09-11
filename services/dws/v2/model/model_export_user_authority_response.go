package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportUserAuthorityResponse Response Object
type ExportUserAuthorityResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ExportUserAuthorityResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportUserAuthorityResponse struct{}"
	}

	return strings.Join([]string{"ExportUserAuthorityResponse", string(data)}, " ")
}
