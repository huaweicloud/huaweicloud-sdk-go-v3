package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowConfigSetRequest Request Object
type ShowConfigSetRequest struct {

	// 配置集合id
	Configsetid string `json:"configsetid"`
}

func (o ShowConfigSetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowConfigSetRequest struct{}"
	}

	return strings.Join([]string{"ShowConfigSetRequest", string(data)}, " ")
}
