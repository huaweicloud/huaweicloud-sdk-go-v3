package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportIpBlacklistResponse Response Object
type ExportIpBlacklistResponse struct {
	Body *interface{} `json:"body,omitempty"`

	ContentDisposition *string `json:"Content-Disposition,omitempty"`

	ContentLength *int32 `json:"Content-Length,omitempty"`

	ContentType    *string `json:"Content-Type,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ExportIpBlacklistResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportIpBlacklistResponse struct{}"
	}

	return strings.Join([]string{"ExportIpBlacklistResponse", string(data)}, " ")
}
