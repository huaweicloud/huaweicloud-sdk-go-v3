package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ModifyTaskInfo Information of task
type ModifyTaskInfo struct {

	// **参数解释**: 更新的待办动作 - TERMINATE 终止待办 - CONTINUE 继续执行 - ADD_COMMENT 添加评论 - DELETE_COMMENT 删除评论  - ADD_ATTACHMENT 添加附件 - DELETE_ATTACHMENT 删除附件   **约束限制**: - TERMINATE 不涉及 - CONTINUE 不涉及 - ADD_COMMENT 需要和请求参数comment配合使用 - DELETE_COMMENT 需要和请求参数comment_id配合使用  - ADD_ATTACHMENT 需要和请求参数attachment_id配合使用 - DELETE_ATTACHMENT 需要和请求参数attachment_id配合使用    **取值范围**: - TERMINATE - CONTINUE - ADD_COMMENT - DELETE_COMMENT  - ADD_ATTACHMENT - DELETE_ATTACHMENT  **默认值**:  ADD_ATTACHMENT 添加评论
	Action *ModifyTaskInfoAction `json:"action,omitempty"`

	// 附件id
	AttachmentId *string `json:"attachment_id,omitempty"`

	// 待办评论内容
	Comment *string `json:"comment,omitempty"`

	// 待办评论id
	CommentId *string `json:"comment_id,omitempty"`
}

func (o ModifyTaskInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyTaskInfo struct{}"
	}

	return strings.Join([]string{"ModifyTaskInfo", string(data)}, " ")
}

type ModifyTaskInfoAction struct {
	value string
}

type ModifyTaskInfoActionEnum struct {
	TERMINATE        ModifyTaskInfoAction
	CONTINUE         ModifyTaskInfoAction
	ADD_COMMEN       ModifyTaskInfoAction
	DELETE_COMMENT   ModifyTaskInfoAction
	ADD_ATTACHMENT   ModifyTaskInfoAction
	DELETE_ATTACHMEN ModifyTaskInfoAction
}

func GetModifyTaskInfoActionEnum() ModifyTaskInfoActionEnum {
	return ModifyTaskInfoActionEnum{
		TERMINATE: ModifyTaskInfoAction{
			value: "TERMINATE",
		},
		CONTINUE: ModifyTaskInfoAction{
			value: "CONTINUE",
		},
		ADD_COMMEN: ModifyTaskInfoAction{
			value: "ADD_COMMEN",
		},
		DELETE_COMMENT: ModifyTaskInfoAction{
			value: "DELETE_COMMENT",
		},
		ADD_ATTACHMENT: ModifyTaskInfoAction{
			value: "ADD_ATTACHMENT",
		},
		DELETE_ATTACHMEN: ModifyTaskInfoAction{
			value: "DELETE_ATTACHMEN",
		},
	}
}

func (c ModifyTaskInfoAction) Value() string {
	return c.value
}

func (c ModifyTaskInfoAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ModifyTaskInfoAction) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
