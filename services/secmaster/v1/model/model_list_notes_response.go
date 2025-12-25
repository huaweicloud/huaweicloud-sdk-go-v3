package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListNotesResponse Response Object
type ListNotesResponse struct {

	// **参数解释**: 错误码 **取值范围**: 不涉及
	Code *string `json:"code,omitempty"`

	// **参数解释**: 错误描述 **取值范围**: 不涉及
	Message *string `json:"message,omitempty"`

	// 总条数
	Total *int32 `json:"total,omitempty"`

	// 分页查询数据大小
	Size *int32 `json:"size,omitempty"`

	// 当前页码
	Page *int32 `json:"page,omitempty"`

	// 是否请求成功
	Success *bool `json:"success,omitempty"`

	// 评论列表详情
	Data           *[]NotesDetail `json:"data,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ListNotesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListNotesResponse struct{}"
	}

	return strings.Join([]string{"ListNotesResponse", string(data)}, " ")
}
