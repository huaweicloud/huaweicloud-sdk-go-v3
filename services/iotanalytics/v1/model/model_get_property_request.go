package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 查询属性值请求
type GetPropertyRequest struct {
	// 对property按指定tags标签进行过滤查询，填入设备标签与标签值，不可为空，例如 {\"deviceId\": \"id0001\"}

	Tags map[string]string `json:"tags"`
	// 查询设备的属性名称

	PropertyNames *[]string `json:"property_names,omitempty"`
}

func (o GetPropertyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetPropertyRequest struct{}"
	}

	return strings.Join([]string{"GetPropertyRequest", string(data)}, " ")
}
