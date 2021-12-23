package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListAclStrategiesV2Response struct {
	// 本次返回的列表长度

	Size int32 `json:"size"`
	// 满足条件的记录数

	Total int64 `json:"total"`
	// 本次查询返回的ACL策略列表

	Acls           *[]ApiAclInfoWithBindNum `json:"acls,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o ListAclStrategiesV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAclStrategiesV2Response struct{}"
	}

	return strings.Join([]string{"ListAclStrategiesV2Response", string(data)}, " ")
}
