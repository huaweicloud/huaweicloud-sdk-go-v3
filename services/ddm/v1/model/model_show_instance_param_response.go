package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowInstanceParamResponse Response Object
type ShowInstanceParamResponse struct {

	// 参数组id。
	Id *string `json:"id,omitempty"`

	// 参数组名称。
	Name *string `json:"name,omitempty"`

	// 参数组描述。
	Description *string `json:"description,omitempty"`

	// DDM参数最后更新时间。
	Updated *string `json:"updated,omitempty"`

	// DDM实例参数信息列表的集合。
	ConfigurationParameter *[]ConfigurationParameterList `json:"configuration_parameter,omitempty"`

	// 分页参数: 起始值。
	Offset *int32 `json:"offset,omitempty"`

	// 分页参数：每页多少条。
	Limit *int32 `json:"limit,omitempty"`

	// 集合总数
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowInstanceParamResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceParamResponse struct{}"
	}

	return strings.Join([]string{"ShowInstanceParamResponse", string(data)}, " ")
}
