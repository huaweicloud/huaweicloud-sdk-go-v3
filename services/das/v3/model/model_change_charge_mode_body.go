package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ChangeChargeModeBody struct {

	// 实例ID列表
	InstanceIdList []string `json:"instance_id_list"`

	// 引擎类型
	DatastoreType string `json:"datastore_type"`
}

func (o ChangeChargeModeBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeChargeModeBody struct{}"
	}

	return strings.Join([]string{"ChangeChargeModeBody", string(data)}, " ")
}
