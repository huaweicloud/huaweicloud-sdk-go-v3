package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CommentLogInfo 评论记录。
type CommentLogInfo struct {

	// 操作时间,格式遵循：RFC 3339。 例 “2020-07-30T10:43:17Z”
	OperateTime *string `json:"operate_time,omitempty"`

	// * USER_REJECT：用户驳回 * USER_COMMENT：用户评论 * ADMIN_COMMENT：管理员评论
	CommentType *CommentLogInfoCommentType `json:"comment_type,omitempty"`

	// 标题。
	CommentTitle *string `json:"comment_title,omitempty"`

	// 消息。
	CommentMessage *string `json:"comment_message,omitempty"`

	// 附件下载地址
	AttachmentDownloadUrl *[]string `json:"attachment_download_url,omitempty"`
}

func (o CommentLogInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CommentLogInfo struct{}"
	}

	return strings.Join([]string{"CommentLogInfo", string(data)}, " ")
}

type CommentLogInfoCommentType struct {
	value string
}

type CommentLogInfoCommentTypeEnum struct {
	USER_REJECT   CommentLogInfoCommentType
	USER_COMMENT  CommentLogInfoCommentType
	ADMIN_COMMENT CommentLogInfoCommentType
}

func GetCommentLogInfoCommentTypeEnum() CommentLogInfoCommentTypeEnum {
	return CommentLogInfoCommentTypeEnum{
		USER_REJECT: CommentLogInfoCommentType{
			value: "USER_REJECT",
		},
		USER_COMMENT: CommentLogInfoCommentType{
			value: "USER_COMMENT",
		},
		ADMIN_COMMENT: CommentLogInfoCommentType{
			value: "ADMIN_COMMENT",
		},
	}
}

func (c CommentLogInfoCommentType) Value() string {
	return c.value
}

func (c CommentLogInfoCommentType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CommentLogInfoCommentType) UnmarshalJSON(b []byte) error {
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
