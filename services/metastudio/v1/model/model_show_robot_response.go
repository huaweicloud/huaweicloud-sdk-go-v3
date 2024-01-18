package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRobotResponse Response Object
type ShowRobotResponse struct {

	// 应用ID。
	RobotId *string `json:"robot_id,omitempty"`

	// 应用名称。
	Name *string `json:"name,omitempty"`

	// 第三方应用ID。
	AppId *string `json:"app_id,omitempty"`

	// 对接第三方应用厂商类型。 > 0：科大讯飞AIUI；1：华为云CBS；2：科大讯飞星火交互认知大模型；5：第三方驱动
	AppType *int32 `json:"app_type,omitempty"`

	// 对话的并发数
	Concurrency *int32 `json:"concurrency,omitempty"`

	// 创建时间，格式遵循：RFC 3339 如\"2021-01-10T08:43:17Z\"。
	CreateTime *string `json:"create_time,omitempty"`

	// 更新时间，格式遵循：RFC 3339 如\"2021-01-10T08:43:17Z\"。
	UpdateTime *string `json:"update_time,omitempty"`

	// CBS所在区域
	Region *int32 `json:"region,omitempty"`

	// CBS所在区域的projectId
	CbsProjectId *string `json:"cbs_project_id,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowRobotResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRobotResponse struct{}"
	}

	return strings.Join([]string{"ShowRobotResponse", string(data)}, " ")
}
