package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAppConfigsV2Response Response Object
type ListAppConfigsV2Response struct {

	// 本次返回的列表长度
	Size int32 `json:"size"`

	// 满足条件的记录数
	Total int64 `json:"total"`

	// 本次查询到的应用配置列表
	Configs        *[]AppConfigInfoV2 `json:"configs,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ListAppConfigsV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAppConfigsV2Response struct{}"
	}

	return strings.Join([]string{"ListAppConfigsV2Response", string(data)}, " ")
}
