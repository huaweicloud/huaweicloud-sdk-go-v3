package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TargetRootDb 数据库对象迁移或同步目标库信息，两层到三层数据库时需要指定。
type TargetRootDb struct {

	// 库名。
	DbName *string `json:"db_name,omitempty"`

	// 默认编码格式是utf8。
	DbEncoding *string `json:"db_encoding,omitempty"`
}

func (o TargetRootDb) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TargetRootDb struct{}"
	}

	return strings.Join([]string{"TargetRootDb", string(data)}, " ")
}
