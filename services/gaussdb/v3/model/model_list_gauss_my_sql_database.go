package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 数据库信息。
type ListGaussMySqlDatabase struct {

	// 数据库名称。
	Name *string `json:"name,omitempty" xml:"name"`

	// 是否为只读权限：true表示只读,false表示可读写。
	Readonly *bool `json:"readonly,omitempty" xml:"readonly"`
}

func (o ListGaussMySqlDatabase) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGaussMySqlDatabase struct{}"
	}

	return strings.Join([]string{"ListGaussMySqlDatabase", string(data)}, " ")
}
