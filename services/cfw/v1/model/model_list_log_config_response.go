package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListLogConfigResponse Response Object
type ListLogConfigResponse struct {
	Data           *LogConfigDto `json:"data,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ListLogConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLogConfigResponse struct{}"
	}

	return strings.Join([]string{"ListLogConfigResponse", string(data)}, " ")
}
