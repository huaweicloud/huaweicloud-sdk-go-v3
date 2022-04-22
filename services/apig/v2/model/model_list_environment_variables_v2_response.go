package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListEnvironmentVariablesV2Response struct {

	// 本次返回的列表长度
	Size int32 `json:"size"`

	// 满足条件的记录数
	Total int64 `json:"total"`

	// 本次返回的环境变量列表
	Variables      *[]EnvVariableInfo `json:"variables,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ListEnvironmentVariablesV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEnvironmentVariablesV2Response struct{}"
	}

	return strings.Join([]string{"ListEnvironmentVariablesV2Response", string(data)}, " ")
}
