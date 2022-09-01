package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListAppCodesV2Response struct {

	// 本次返回的列表长度
	Size int32 `json:"size" xml:"size"`

	// 满足条件的记录数
	Total int64 `json:"total" xml:"total"`

	// App Code列表
	AppCodes       *[]AppCodeBaseInfo `json:"app_codes,omitempty" xml:"app_codes"`
	HttpStatusCode int                `json:"-"`
}

func (o ListAppCodesV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAppCodesV2Response struct{}"
	}

	return strings.Join([]string{"ListAppCodesV2Response", string(data)}, " ")
}
