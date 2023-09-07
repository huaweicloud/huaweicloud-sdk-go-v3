package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// QueryColumnReq 是否从源库获取最新的列信息
type QueryColumnReq struct {

	// 是否从Node获取最新的列信息
	IsForceRefresh *bool `json:"is_force_refresh,omitempty"`

	// 指定数据库表信息
	DbObjectInfos []SelectDbTableObjectInfo `json:"db_object_infos"`
}

func (o QueryColumnReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryColumnReq struct{}"
	}

	return strings.Join([]string{"QueryColumnReq", string(data)}, " ")
}
