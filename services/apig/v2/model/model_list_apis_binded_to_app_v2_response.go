package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListApisBindedToAppV2Response struct {

	// 本次返回的列表长度
	Size int32 `json:"size" xml:"size"`

	// 满足条件的记录数
	Total int64 `json:"total" xml:"total"`

	// 本次返回的API列表
	Auths          *[]ApiAuthInfo `json:"auths,omitempty" xml:"auths"`
	HttpStatusCode int            `json:"-"`
}

func (o ListApisBindedToAppV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListApisBindedToAppV2Response struct{}"
	}

	return strings.Join([]string{"ListApisBindedToAppV2Response", string(data)}, " ")
}
