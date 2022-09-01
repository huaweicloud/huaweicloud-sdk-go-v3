package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowAppVersionDetailRequest struct {

	// 铂金版实例ID，专业版实例为空值
	IefInstanceId *string `json:"ief-instance-id,omitempty" xml:"ief-instance-id"`

	// 应用模板ID
	AppId string `json:"app_id" xml:"app_id"`

	// 应用模板版本ID
	VersionId string `json:"version_id" xml:"version_id"`
}

func (o ShowAppVersionDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAppVersionDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowAppVersionDetailRequest", string(data)}, " ")
}
