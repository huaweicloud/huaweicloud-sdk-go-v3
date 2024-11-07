package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchUpgradeDatabaseVersionRequestBody struct {

	// 指定需要升级数据库补丁版本的实例ID列表。 一次最多可传入10个实例ID，且为同一引擎的实例。
	InstanceIds []string `json:"instance_ids"`
}

func (o BatchUpgradeDatabaseVersionRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpgradeDatabaseVersionRequestBody struct{}"
	}

	return strings.Join([]string{"BatchUpgradeDatabaseVersionRequestBody", string(data)}, " ")
}
