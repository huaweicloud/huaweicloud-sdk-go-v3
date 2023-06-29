package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListIaConfigsResponse Response Object
type ListIaConfigsResponse struct {

	// 南向3rdIA配置项列表
	Configs        *[]QueryIaConfigResponseDto `json:"configs,omitempty"`
	HttpStatusCode int                         `json:"-"`
}

func (o ListIaConfigsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListIaConfigsResponse struct{}"
	}

	return strings.Join([]string{"ListIaConfigsResponse", string(data)}, " ")
}
