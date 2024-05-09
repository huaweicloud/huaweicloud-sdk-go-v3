package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListStarRocksDataReplicationConfigRequest Request Object
type ListStarRocksDataReplicationConfigRequest struct {

	// StarRocks实例ID，严格匹配UUID规则。
	InstanceId string `json:"instance_id"`

	// 请求语言类型。默认en-us。 取值范围： - en-us - zh-cn
	XLanguage string `json:"X-Language"`

	// 数据同步任务名。 字符长度限制3~128位，仅支持英文大小写字母、数字以及下划线_。
	TaskName string `json:"task_name"`
}

func (o ListStarRocksDataReplicationConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListStarRocksDataReplicationConfigRequest struct{}"
	}

	return strings.Join([]string{"ListStarRocksDataReplicationConfigRequest", string(data)}, " ")
}
