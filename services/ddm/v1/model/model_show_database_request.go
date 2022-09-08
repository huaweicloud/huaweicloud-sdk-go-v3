package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowDatabaseRequest struct {

	// DDM实例ID
	InstanceId string `json:"instance_id"`

	// 需要查询的逻辑库名称，不区分大小写。
	DdmDbname string `json:"ddm_dbname"`
}

func (o ShowDatabaseRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDatabaseRequest struct{}"
	}

	return strings.Join([]string{"ShowDatabaseRequest", string(data)}, " ")
}
