package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAccessTopResponse Response Object
type ShowAccessTopResponse struct {
	Data           *AccessTopVo `json:"data,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ShowAccessTopResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAccessTopResponse struct{}"
	}

	return strings.Join([]string{"ShowAccessTopResponse", string(data)}, " ")
}
