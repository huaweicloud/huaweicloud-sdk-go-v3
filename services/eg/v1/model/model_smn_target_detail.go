package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SmnTargetDetail 订阅的SMN事件目标参数列表，该字段序列化后总长度不超过1024字节
type SmnTargetDetail struct {

	// 主题urn
	Urn string `json:"urn"`

	// 委托名称
	AgencyName string `json:"agency_name"`

	SubjectTransform *SmnTargetDetailSubjectTransform `json:"subject_transform,omitempty"`
}

func (o SmnTargetDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SmnTargetDetail struct{}"
	}

	return strings.Join([]string{"SmnTargetDetail", string(data)}, " ")
}
