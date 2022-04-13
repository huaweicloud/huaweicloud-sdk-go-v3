package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListAppCodesV2Response struct {
	// 本次返回的列表长度

	Size int32 `json:"size"`
	// 满足条件的记录数

	Total int64 `json:"total"`
	// App Code列表

	AppCodes       *[]AppCodeBaseInfo `json:"app_codes,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ListAppCodesV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAppCodesV2Response struct{}"
	}

	return strings.Join([]string{"ListAppCodesV2Response", string(data)}, " ")
}
