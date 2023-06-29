package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StudyRsp study详情返回体
type StudyRsp struct {

	// study名称
	Name *string `json:"name,omitempty"`

	// study id
	Id *string `json:"id,omitempty"`

	// 医疗项目名称
	EihealthProjectName *string `json:"eihealth_project_name,omitempty"`

	// 医疗项目id
	EihealthProjectId *string `json:"eihealth_project_id,omitempty"`

	// study描述
	Description *string `json:"description,omitempty"`

	// study创建时间
	CreateTime *string `json:"create_time,omitempty"`

	// study更新时间
	UpdateTime *string `json:"update_time,omitempty"`

	LatestJob *StudyJobRsp `json:"latest_job,omitempty"`
}

func (o StudyRsp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StudyRsp struct{}"
	}

	return strings.Join([]string{"StudyRsp", string(data)}, " ")
}
