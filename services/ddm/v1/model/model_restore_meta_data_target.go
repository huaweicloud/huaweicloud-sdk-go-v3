package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RestoreMetaDataTarget struct {

	// metadata恢复目标dn。
	DataNodes []string `json:"data_nodes"`

	// 实例id。
	InstanceId string `json:"instance_id"`
}

func (o RestoreMetaDataTarget) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestoreMetaDataTarget struct{}"
	}

	return strings.Join([]string{"RestoreMetaDataTarget", string(data)}, " ")
}
