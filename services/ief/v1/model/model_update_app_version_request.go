package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateAppVersionRequest struct {

	// 铂金版实例ID，专业版实例为空值
	IefInstanceId *string `json:"ief-instance-id,omitempty" xml:"ief-instance-id"`

	// 应用模板ID
	AppId string `json:"app_id" xml:"app_id"`

	// 应用模板版本ID
	VersionId string `json:"version_id" xml:"version_id"`

	Body *VersionUpdate `json:"body,omitempty" xml:"body"`
}

func (o UpdateAppVersionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAppVersionRequest struct{}"
	}

	return strings.Join([]string{"UpdateAppVersionRequest", string(data)}, " ")
}
