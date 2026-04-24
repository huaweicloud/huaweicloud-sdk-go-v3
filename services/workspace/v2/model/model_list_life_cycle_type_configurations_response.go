package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListLifeCycleTypeConfigurationsResponse Response Object
type ListLifeCycleTypeConfigurationsResponse struct {

	// 生命周期类型配置响应。
	Body           map[string]LifeCycleTypeConfigRsp `json:"body,omitempty"`
	HttpStatusCode int                               `json:"-"`
}

func (o ListLifeCycleTypeConfigurationsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLifeCycleTypeConfigurationsResponse struct{}"
	}

	return strings.Join([]string{"ListLifeCycleTypeConfigurationsResponse", string(data)}, " ")
}
