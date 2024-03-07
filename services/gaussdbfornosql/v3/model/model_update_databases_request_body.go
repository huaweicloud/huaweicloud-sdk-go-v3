package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateDatabasesRequestBody struct {

	// 对实例的操作: - flush:清理数据
	Action string `json:"action"`

	// 指定需要清理的DB_ID,当action为flush时,才会生效。
	DbId *int32 `json:"db_id,omitempty"`
}

func (o UpdateDatabasesRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDatabasesRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateDatabasesRequestBody", string(data)}, " ")
}
