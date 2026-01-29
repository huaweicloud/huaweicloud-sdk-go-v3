package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPreProcessRulesListResponse Response Object
type ShowPreProcessRulesListResponse struct {

	// **参数解释**: 错误码 **取值范围**: 不涉及
	Code *string `json:"code,omitempty"`

	// **参数解释**: 错误描述 **取值范围**: 不涉及
	Message *string `json:"message,omitempty"`

	// 分页查询数据大小
	Size *int32 `json:"size,omitempty"`

	// 当前页码
	Page *int32 `json:"page,omitempty"`

	// 总数
	Total *int32 `json:"total,omitempty"`

	// 分类映射信息集合
	Data           *[]ShowPreProcessRulesListResponseBodyData `json:"data,omitempty"`
	HttpStatusCode int                                        `json:"-"`
}

func (o ShowPreProcessRulesListResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPreProcessRulesListResponse struct{}"
	}

	return strings.Join([]string{"ShowPreProcessRulesListResponse", string(data)}, " ")
}
