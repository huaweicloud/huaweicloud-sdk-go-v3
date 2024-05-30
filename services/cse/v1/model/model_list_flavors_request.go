package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListFlavorsRequest Request Object
type ListFlavorsRequest struct {

	// 微服务引擎应用部署类型，查询ServiceComb引擎专享版需要将该值设置为CSE2，查询注册配置中心需要将该值设置为Nacos2，查询网关需要将该值设置为MicroGateway。
	SpecType *string `json:"spec_type,omitempty"`
}

func (o ListFlavorsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFlavorsRequest struct{}"
	}

	return strings.Join([]string{"ListFlavorsRequest", string(data)}, " ")
}
