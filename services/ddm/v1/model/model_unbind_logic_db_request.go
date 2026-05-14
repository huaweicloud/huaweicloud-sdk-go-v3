package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UnbindLogicDbRequest Request Object
type UnbindLogicDbRequest struct {

	// DDM实例ID
	InstanceId string `json:"instance_id"`

	// 需要解绑的逻辑库名
	LogicDbName string `json:"logic_db_name"`
}

func (o UnbindLogicDbRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UnbindLogicDbRequest struct{}"
	}

	return strings.Join([]string{"UnbindLogicDbRequest", string(data)}, " ")
}
