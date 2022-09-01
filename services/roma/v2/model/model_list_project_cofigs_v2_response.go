package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListProjectCofigsV2Response struct {

	// 本次返回的列表长度
	Size int32 `json:"size" xml:"size"`

	// 满足条件的记录数
	Total int64 `json:"total" xml:"total"`

	// 配额列表
	Configs        *[]Config `json:"configs,omitempty" xml:"configs"`
	HttpStatusCode int       `json:"-"`
}

func (o ListProjectCofigsV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProjectCofigsV2Response struct{}"
	}

	return strings.Join([]string{"ListProjectCofigsV2Response", string(data)}, " ")
}
