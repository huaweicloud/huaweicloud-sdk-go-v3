package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEnvironmentsV2Response Response Object
type ListEnvironmentsV2Response struct {

	// 本次返回的列表长度
	Size int32 `json:"size"`

	// 满足条件的记录数
	Total int64 `json:"total"`

	// 本次返回的环境列表
	Envs           *[]EnvInfo `json:"envs,omitempty"`
	HttpStatusCode int        `json:"-"`
}

func (o ListEnvironmentsV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEnvironmentsV2Response struct{}"
	}

	return strings.Join([]string{"ListEnvironmentsV2Response", string(data)}, " ")
}
