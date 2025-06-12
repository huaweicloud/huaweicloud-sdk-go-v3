package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRecycleBinServerResponse Response Object
type ShowRecycleBinServerResponse struct {
	Server         *ServerDetail `json:"server,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ShowRecycleBinServerResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRecycleBinServerResponse struct{}"
	}

	return strings.Join([]string{"ShowRecycleBinServerResponse", string(data)}, " ")
}
