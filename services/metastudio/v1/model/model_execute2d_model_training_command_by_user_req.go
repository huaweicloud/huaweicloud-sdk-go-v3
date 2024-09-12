package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Execute2dModelTrainingCommandByUserReq 租户执行分身数字人模型训练任务命令请求。
type Execute2dModelTrainingCommandByUserReq struct {

	// 命令类型。 * SUBMITVERIFYING: 提交审核 * CONFIRM_ACCEPT：用户确认训练效果 * CONFIRM_REJECT：用户驳回训练效果 * CONFIRM_ANSWER：用户答复 * CONFIRM_PENDING：用户挂起任务 * CONFIRM_ACTIVE：用户激活任务 * GET_MULTIPART_UPLOADED：获取训练视频分片上传地址 * CONFIRM_REPAIR:用户发起优化模型请求 * GET_MULTIPART_UPLOADED：获取训练视频已上传分片信息 * CONFIRM_MULTIPART_UPLOADED：确认训练视频所有分片文件已上传 * GET_ACTION_VIDEO_MULTIPART_UPLOADED：获取动作编排视频分片上传地址 * CONFIRM_ACTION_VIDEO_MULTIPART_UPLOADED：确认动作编排视频所有分片文件已上传 > * CONFIRM_ACCEPT/CONFIRM_REJECT/CONFIRM_ANSWER/CONFIRM_PENDING/CONFIRM_ACTIVE命令仅NA白名单用户可用。
	Command Execute2dModelTrainingCommandByUserReqCommand `json:"command"`

	// 命令类型： * UPDATE_VIDEO: 更新视频 * UPLOAD_VIDEO: 上传视频 * CONFIRM_ACTION_VIDEO: 确认动作编排视频 * GET_ACTION_VIDEO_MULTIPART: 获取动作编排视频分片
	CommandMessage *Execute2dModelTrainingCommandByUserReqCommandMessage `json:"command_message,omitempty"`

	CommentData *CommentData `json:"comment_data,omitempty"`
}

func (o Execute2dModelTrainingCommandByUserReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Execute2dModelTrainingCommandByUserReq struct{}"
	}

	return strings.Join([]string{"Execute2dModelTrainingCommandByUserReq", string(data)}, " ")
}

type Execute2dModelTrainingCommandByUserReqCommand struct {
	value string
}

type Execute2dModelTrainingCommandByUserReqCommandEnum struct {
	SUBMITVERIFYING                         Execute2dModelTrainingCommandByUserReqCommand
	CONFIRM_ACCEPT                          Execute2dModelTrainingCommandByUserReqCommand
	CONFIRM_REJECT                          Execute2dModelTrainingCommandByUserReqCommand
	CONFIRM_ANSWER                          Execute2dModelTrainingCommandByUserReqCommand
	CONFIRM_PENDING                         Execute2dModelTrainingCommandByUserReqCommand
	CONFIRM_ACTIVE                          Execute2dModelTrainingCommandByUserReqCommand
	GET_MULTIPART_UPLOADED                  Execute2dModelTrainingCommandByUserReqCommand
	CONFIRM_REPAIR                          Execute2dModelTrainingCommandByUserReqCommand
	CONFIRM_MULTIPART_UPLOADED              Execute2dModelTrainingCommandByUserReqCommand
	GET_ACTION_VIDEO_MULTIPART_UPLOADED     Execute2dModelTrainingCommandByUserReqCommand
	CONFIRM_ACTION_VIDEO_MULTIPART_UPLOADED Execute2dModelTrainingCommandByUserReqCommand
}

func GetExecute2dModelTrainingCommandByUserReqCommandEnum() Execute2dModelTrainingCommandByUserReqCommandEnum {
	return Execute2dModelTrainingCommandByUserReqCommandEnum{
		SUBMITVERIFYING: Execute2dModelTrainingCommandByUserReqCommand{
			value: "SUBMITVERIFYING",
		},
		CONFIRM_ACCEPT: Execute2dModelTrainingCommandByUserReqCommand{
			value: "CONFIRM_ACCEPT",
		},
		CONFIRM_REJECT: Execute2dModelTrainingCommandByUserReqCommand{
			value: "CONFIRM_REJECT",
		},
		CONFIRM_ANSWER: Execute2dModelTrainingCommandByUserReqCommand{
			value: "CONFIRM_ANSWER",
		},
		CONFIRM_PENDING: Execute2dModelTrainingCommandByUserReqCommand{
			value: "CONFIRM_PENDING",
		},
		CONFIRM_ACTIVE: Execute2dModelTrainingCommandByUserReqCommand{
			value: "CONFIRM_ACTIVE",
		},
		GET_MULTIPART_UPLOADED: Execute2dModelTrainingCommandByUserReqCommand{
			value: "GET_MULTIPART_UPLOADED",
		},
		CONFIRM_REPAIR: Execute2dModelTrainingCommandByUserReqCommand{
			value: "CONFIRM_REPAIR",
		},
		CONFIRM_MULTIPART_UPLOADED: Execute2dModelTrainingCommandByUserReqCommand{
			value: "CONFIRM_MULTIPART_UPLOADED",
		},
		GET_ACTION_VIDEO_MULTIPART_UPLOADED: Execute2dModelTrainingCommandByUserReqCommand{
			value: "GET_ACTION_VIDEO_MULTIPART_UPLOADED",
		},
		CONFIRM_ACTION_VIDEO_MULTIPART_UPLOADED: Execute2dModelTrainingCommandByUserReqCommand{
			value: "CONFIRM_ACTION_VIDEO_MULTIPART_UPLOADED",
		},
	}
}

func (c Execute2dModelTrainingCommandByUserReqCommand) Value() string {
	return c.value
}

func (c Execute2dModelTrainingCommandByUserReqCommand) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *Execute2dModelTrainingCommandByUserReqCommand) UnmarshalJSON(b []byte) error {
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

type Execute2dModelTrainingCommandByUserReqCommandMessage struct {
	value string
}

type Execute2dModelTrainingCommandByUserReqCommandMessageEnum struct {
	UPDATE_VIDEO               Execute2dModelTrainingCommandByUserReqCommandMessage
	UPLOAD_VIDEO               Execute2dModelTrainingCommandByUserReqCommandMessage
	CONFIRM_ACTION_VIDEO       Execute2dModelTrainingCommandByUserReqCommandMessage
	GET_ACTION_VIDEO_MULTIPART Execute2dModelTrainingCommandByUserReqCommandMessage
}

func GetExecute2dModelTrainingCommandByUserReqCommandMessageEnum() Execute2dModelTrainingCommandByUserReqCommandMessageEnum {
	return Execute2dModelTrainingCommandByUserReqCommandMessageEnum{
		UPDATE_VIDEO: Execute2dModelTrainingCommandByUserReqCommandMessage{
			value: "UPDATE_VIDEO",
		},
		UPLOAD_VIDEO: Execute2dModelTrainingCommandByUserReqCommandMessage{
			value: "UPLOAD_VIDEO",
		},
		CONFIRM_ACTION_VIDEO: Execute2dModelTrainingCommandByUserReqCommandMessage{
			value: "CONFIRM_ACTION_VIDEO",
		},
		GET_ACTION_VIDEO_MULTIPART: Execute2dModelTrainingCommandByUserReqCommandMessage{
			value: "GET_ACTION_VIDEO_MULTIPART",
		},
	}
}

func (c Execute2dModelTrainingCommandByUserReqCommandMessage) Value() string {
	return c.value
}

func (c Execute2dModelTrainingCommandByUserReqCommandMessage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *Execute2dModelTrainingCommandByUserReqCommandMessage) UnmarshalJSON(b []byte) error {
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
