package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListDuplicateApisForAppV2Response struct {
	// 本次返回的列表长度

	Size int32 `json:"size"`
	// 满足条件的记录数

	Total int64 `json:"total"`
	// 应用下所有路径冲突的api信息列表

	Apis           *[]ApiDuplicationInfo `json:"apis,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o ListDuplicateApisForAppV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDuplicateApisForAppV2Response struct{}"
	}

	return strings.Join([]string{"ListDuplicateApisForAppV2Response", string(data)}, " ")
}
