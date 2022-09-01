package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListInstancesV2Response struct {

	// 本次返回的列表长度
	Size int32 `json:"size" xml:"size"`

	// 满足条件的记录数
	Total int64 `json:"total" xml:"total"`

	// 实例列表
	Instances      *[]RespInstanceBase `json:"instances,omitempty" xml:"instances"`
	HttpStatusCode int                 `json:"-"`
}

func (o ListInstancesV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstancesV2Response struct{}"
	}

	return strings.Join([]string{"ListInstancesV2Response", string(data)}, " ")
}
