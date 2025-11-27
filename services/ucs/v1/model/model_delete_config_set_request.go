package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteConfigSetRequest Request Object
type DeleteConfigSetRequest struct {

	// 配置集合id
	Configsetid string `json:"configsetid"`
}

func (o DeleteConfigSetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteConfigSetRequest struct{}"
	}

	return strings.Join([]string{"DeleteConfigSetRequest", string(data)}, " ")
}
