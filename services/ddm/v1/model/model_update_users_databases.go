package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateUsersDatabases struct {

	// 逻辑库名称，不区分大小写，databases和name字段必须同时缺失或者同时存在。  默认值为空
	Name *string `json:"name,omitempty"`
}

func (o UpdateUsersDatabases) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateUsersDatabases struct{}"
	}

	return strings.Join([]string{"UpdateUsersDatabases", string(data)}, " ")
}
