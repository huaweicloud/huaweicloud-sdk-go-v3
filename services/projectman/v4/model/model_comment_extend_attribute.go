package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CommentExtendAttribute struct {
	Operator *UserVo `json:"operator,omitempty"`

	// 操作人Id
	OperatorId *string `json:"operator_id,omitempty"`

	// 系统生成评论时执行的动作
	Action *string `json:"action,omitempty"`

	// 系统生成评论时执行的动作(英文)
	ActionUs *string `json:"action_us,omitempty"`

	// 系统生成评论对应的对象类型
	ObjectType *string `json:"object_type,omitempty"`

	// 工作项流转前的状态Code
	PreStatusCode *string `json:"pre_status_code,omitempty"`

	// 工作项流转后的状态Code
	NewStatusCode *string `json:"new_status_code,omitempty"`

	// 对象类型根据field_type_id值变化而变化。 field_type_id=10001时，为StatusVO field_type_id=10007、10008时，为字符串 field_type_id=10003、10004时，为日期时间 field_type_id=10005、10006时，为数字
	PreStatus *interface{} `json:"pre_status,omitempty"`

	// 对象类型根据field_type_id值变化而变化。 field_type_id=10001、10002时，为StatusVO field_type_id=10007、10008时，为字符串 field_type_id=10003、10004时，为日期时间 field_type_id=10005、10006时，为数字 field_type_id=10009、10010时，为UserVO
	NewStatus *interface{} `json:"new_status,omitempty"`

	// 字段类型
	FieldType *string `json:"field_type,omitempty"`

	// 字段类型对应的Id
	FieldTypeId *string `json:"field_type_id,omitempty"`

	// 字段显示名
	DisplayName *string `json:"display_name,omitempty"`
}

func (o CommentExtendAttribute) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CommentExtendAttribute struct{}"
	}

	return strings.Join([]string{"CommentExtendAttribute", string(data)}, " ")
}
