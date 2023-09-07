package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// QuerySelectObjectInfoReq 数据库对象采集请求体
type QuerySelectObjectInfoReq struct {

	// 查询指定库的信息。
	DbNames *[]string `json:"db_names,omitempty"`

	// 查询对象信息类型，取值： - source：查询源库对象信息。 - modified：查询已选择的（已同步的和未下发的）对象信息。 - synchronized：查询已同步的（已下发的）对象信息 ， 使用场景在任务处于全量中或者增量中。
	Type string `json:"type"`

	// 是否强制刷新。取值： - true：是，表示从源库重新查询。 - false：否，表示从已缓存中数据查询。
	IsRefresh *string `json:"is_refresh,omitempty"`
}

func (o QuerySelectObjectInfoReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QuerySelectObjectInfoReq struct{}"
	}

	return strings.Join([]string{"QuerySelectObjectInfoReq", string(data)}, " ")
}
