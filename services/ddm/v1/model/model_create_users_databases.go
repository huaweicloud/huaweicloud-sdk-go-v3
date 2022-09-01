package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateUsersDatabases struct {

	// 关联逻辑库名称。
	Name string `json:"name" xml:"name"`
}

func (o CreateUsersDatabases) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateUsersDatabases struct{}"
	}

	return strings.Join([]string{"CreateUsersDatabases", string(data)}, " ")
}
