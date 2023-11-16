package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteBasicPluginRequest Request Object
type DeleteBasicPluginRequest struct {

	// 租户ID
	DomainId string `json:"domain_id"`

	// 是否调用成功
	PluginName *string `json:"plugin_name,omitempty"`
}

func (o DeleteBasicPluginRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteBasicPluginRequest struct{}"
	}

	return strings.Join([]string{"DeleteBasicPluginRequest", string(data)}, " ")
}
