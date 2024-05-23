package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HuaweiEiCbs 华为云CBS应用配置
type HuaweiEiCbs struct {

	// CBS应用ID。
	AppId *string `json:"app_id,omitempty"`

	// CBS所在区域
	Region *int32 `json:"region,omitempty"`

	// CBS所在区域的projectId
	CbsProjectId *string `json:"cbs_project_id,omitempty"`

	// SIS所在区域
	SisRegion *int32 `json:"sis_region,omitempty"`

	// SIS所在区域的projectId
	SisProjectId *string `json:"sis_project_id,omitempty"`
}

func (o HuaweiEiCbs) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HuaweiEiCbs struct{}"
	}

	return strings.Join([]string{"HuaweiEiCbs", string(data)}, " ")
}
