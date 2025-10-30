package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchUpgradeDatabasesRequestBody struct {

	// 要版本升级的批量实例。
	DatabasesInstanceInfos []UpgradeDatabasesSingleInstance `json:"databases_instance_infos"`

	// 是否延迟升级。
	Delay string `json:"delay"`
}

func (o BatchUpgradeDatabasesRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpgradeDatabasesRequestBody struct{}"
	}

	return strings.Join([]string{"BatchUpgradeDatabasesRequestBody", string(data)}, " ")
}
