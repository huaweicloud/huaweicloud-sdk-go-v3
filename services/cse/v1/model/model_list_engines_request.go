package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEnginesRequest Request Object
type ListEnginesRequest struct {

	// 偏移量。
	Offset *int32 `json:"offset,omitempty"`

	// 每页显示的条目数量。
	Limit *string `json:"limit,omitempty"`

	// 查询所有微服务引擎需要将该值设置为ALL，查询ServiceComb引擎专享版需要将该值设置为CSE，查询注册配置中心需要将该值设置为Nacos，查询网关需要将该值设置为MicroGateway。
	Type *string `json:"type,omitempty"`
}

func (o ListEnginesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEnginesRequest struct{}"
	}

	return strings.Join([]string{"ListEnginesRequest", string(data)}, " ")
}
