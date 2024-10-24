package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDdmDatabaseRequest Request Object
type DeleteDdmDatabaseRequest struct {

	// DDM实例ID。
	InstanceId string `json:"instance_id"`

	// 逻辑库名称。
	DatabaseName string `json:"database_name"`

	// 是否同时删除关联后端数据库实例上存储的数据。 - 取值为true：删除。 - 取值为false：不删除。
	DeleteDnData bool `json:"delete_dn_data"`
}

func (o DeleteDdmDatabaseRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDdmDatabaseRequest struct{}"
	}

	return strings.Join([]string{"DeleteDdmDatabaseRequest", string(data)}, " ")
}
