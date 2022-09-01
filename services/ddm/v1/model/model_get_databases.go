package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// databases返回参数
type GetDatabases struct {

	// 分片数。
	Dbslot int32 `json:"dbslot" xml:"dbslot"`

	// 分片名称.
	Name string `json:"name" xml:"name"`

	// 状态。
	Status string `json:"status" xml:"status"`

	// 创建时间。
	Created string `json:"created" xml:"created"`

	// 最近更新时间。
	Updated string `json:"updated" xml:"updated"`

	// 所在RDS的id。
	Id string `json:"id" xml:"id"`

	// 所在RDS的名称。
	IdName string `json:"idName" xml:"idName"`
}

func (o GetDatabases) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetDatabases struct{}"
	}

	return strings.Join([]string{"GetDatabases", string(data)}, " ")
}
