package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ServerIdBean ServerIdBean
type ServerIdBean struct {

	// 实例ID。可在查询实例列表接口的ID字段获取。
	InstanceId string `json:"instance_id"`
}

func (o ServerIdBean) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServerIdBean struct{}"
	}

	return strings.Join([]string{"ServerIdBean", string(data)}, " ")
}
