package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListConfigMapsResponse struct {

	// 配置项
	Configmaps *[]ConfigMapResp `json:"configmaps,omitempty" xml:"configmaps"`

	// 满足条件的个数
	Count          *int32 `json:"count,omitempty" xml:"count"`
	HttpStatusCode int    `json:"-"`
}

func (o ListConfigMapsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListConfigMapsResponse struct{}"
	}

	return strings.Join([]string{"ListConfigMapsResponse", string(data)}, " ")
}
