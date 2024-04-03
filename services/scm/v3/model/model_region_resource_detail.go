package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RegionResourceDetail struct {

	// 局点ID。当服务为全局服务时，region_id为global，其余按照IAM的标准命名。
	RegionId string `json:"region_id"`

	// 请求当前region资源信息过程中，响应是否存在异常的标志。 - true : 存在异常，当前region所统计数据不准确 - false: 无异常，当前region所统计数据准确
	IsError bool `json:"is_error"`

	// 资源集合，每个资源的标识：资源ID + “:” + 资源名称，详情请参见ResourceDetail字段数据结构说明。
	Resources []ResourceDetail `json:"resources"`
}

func (o RegionResourceDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RegionResourceDetail struct{}"
	}

	return strings.Join([]string{"RegionResourceDetail", string(data)}, " ")
}
