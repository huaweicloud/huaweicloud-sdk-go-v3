package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDdmInstanceRequest Request Object
type DeleteDdmInstanceRequest struct {

	// DDM实例ID。
	InstanceId string `json:"instance_id"`

	// 是否同时删除关联后端数据库实例上存储的数据。  - 取值true：删除。 - 取值为空或false：不删除。 默认值为空。
	DeleteDnData *bool `json:"delete_dn_data,omitempty"`
}

func (o DeleteDdmInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDdmInstanceRequest struct{}"
	}

	return strings.Join([]string{"DeleteDdmInstanceRequest", string(data)}, " ")
}
