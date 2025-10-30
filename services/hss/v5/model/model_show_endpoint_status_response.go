package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowEndpointStatusResponse Response Object
type ShowEndpointStatusResponse struct {

	// **参数解释** 服务器状态的列表 **约束限制**: 不涉及
	DataList       *[]HostStatusInfo `json:"data_list,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ShowEndpointStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowEndpointStatusResponse struct{}"
	}

	return strings.Join([]string{"ShowEndpointStatusResponse", string(data)}, " ")
}
