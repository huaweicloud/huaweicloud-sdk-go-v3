package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListConfigMapsResponse Response Object
type ListConfigMapsResponse struct {

	// 配置项数量
	Count *int32 `json:"count,omitempty"`

	// 配置项详情
	Configmaps     *[]ConfigMap `json:"configmaps,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ListConfigMapsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListConfigMapsResponse struct{}"
	}

	return strings.Join([]string{"ListConfigMapsResponse", string(data)}, " ")
}
