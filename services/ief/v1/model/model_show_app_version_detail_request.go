package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAppVersionDetailRequest Request Object
type ShowAppVersionDetailRequest struct {

	// 铂金版实例ID，专业版实例为空值
	IefInstanceId *string `json:"ief-instance-id,omitempty"`

	// 应用模板ID
	AppId string `json:"app_id"`

	// 应用模板版本ID
	VersionId string `json:"version_id"`
}

func (o ShowAppVersionDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAppVersionDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowAppVersionDetailRequest", string(data)}, " ")
}
