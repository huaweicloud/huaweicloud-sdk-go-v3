package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListApisNotBoundWithSignatureKeyV2Response struct {
	// 本次返回的列表长度

	Size int32 `json:"size"`
	// 满足条件的记录数

	Total int64 `json:"total"`
	// 本次查询返回的API列表

	Apis           *[]ApiForSign `json:"apis,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ListApisNotBoundWithSignatureKeyV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListApisNotBoundWithSignatureKeyV2Response struct{}"
	}

	return strings.Join([]string{"ListApisNotBoundWithSignatureKeyV2Response", string(data)}, " ")
}
