package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ListRtcUserListRequest struct {

	// 使用AK/SK方式认证时必选，携带的鉴权信息。
	Authorization *string `json:"Authorization,omitempty"`

	// 使用AK/SK方式认证时必选，请求的发生时间。
	XSdkDate *string `json:"X-Sdk-Date,omitempty"`

	// 使用AK/SK方式认证时必选，携带项目ID信息，与路径参数中的项目ID相同。
	XProjectId *string `json:"X-Project-Id,omitempty"`

	// 应用id
	App string `json:"app"`

	// 房间id
	RoomId *string `json:"room_id,omitempty"`

	// 用户id
	Uid *string `json:"uid,omitempty"`

	// 用户昵称
	Nickname *string `json:"nickname,omitempty"`

	// 用户省份，支持省份名或缩写，如广东或者GD
	Region *[]string `json:"region,omitempty"`

	// 用户接入运营商
	Isp *[]string `json:"isp,omitempty"`

	// 用户状态，取值如下： - FAIL：加入失败 - ONLINE：在线 - OFFLINE：离开
	State *[]string `json:"state,omitempty"`

	// 查询起始时间。UTC时间，格式：YYYY-MM-DDThh:mm:ssZ，如2020-04-23T06:00:00Z，不写默认读取过去1小时数据数据。
	StartTime *string `json:"start_time,omitempty"`

	// 查询结束时间。UTC时间，格式：YYYY-MM-DDThh:mm:ssZ，如2020-04-23T06:00:00Z，不写默认为当前时间。
	EndTime *string `json:"end_time,omitempty"`

	// 查询结果限制
	Limit *int32 `json:"limit,omitempty"`

	// 查询偏移量
	Offset *int32 `json:"offset,omitempty"`

	// 查询模式，取值如下： - detail：会话级 - summary：用户级（默认）
	Type *ListRtcUserListRequestType `json:"type,omitempty"`
}

func (o ListRtcUserListRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRtcUserListRequest struct{}"
	}

	return strings.Join([]string{"ListRtcUserListRequest", string(data)}, " ")
}

type ListRtcUserListRequestType struct {
	value string
}

type ListRtcUserListRequestTypeEnum struct {
	DETAIL  ListRtcUserListRequestType
	SUMMARY ListRtcUserListRequestType
}

func GetListRtcUserListRequestTypeEnum() ListRtcUserListRequestTypeEnum {
	return ListRtcUserListRequestTypeEnum{
		DETAIL: ListRtcUserListRequestType{
			value: "detail",
		},
		SUMMARY: ListRtcUserListRequestType{
			value: "summary",
		},
	}
}

func (c ListRtcUserListRequestType) Value() string {
	return c.value
}

func (c ListRtcUserListRequestType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListRtcUserListRequestType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
