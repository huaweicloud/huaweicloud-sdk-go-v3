package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type LtsAccessConfigInfoRespon200 struct {

	// 跨账号日志接入id
	AccessConfigId *string `json:"access_config_id,omitempty"`

	// 项目ID
	ProjectId *string `json:"project_id,omitempty"`

	// 跨账号日志接入名称
	AccessConfigName *string `json:"access_config_name,omitempty"`

	// 跨账号日志接入类型
	AccessConfigType *interface{} `json:"access_config_type,omitempty"`

	// 日志组ID
	GroupId *string `json:"group_id,omitempty"`

	// 日志组名称
	LogGroupName *string `json:"log_group_name,omitempty"`

	// 日志流ID
	LogStreamId *string `json:"log_stream_id,omitempty"`

	// 日志流名称
	LogStreamName *string `json:"log_stream_name,omitempty"`

	// 创建时间
	CreateTime *int64 `json:"create_time,omitempty"`

	AgencyLogAccess *PreviewAgencyLogAccessReqBody `json:"agency_log_access,omitempty"`
}

func (o LtsAccessConfigInfoRespon200) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LtsAccessConfigInfoRespon200 struct{}"
	}

	return strings.Join([]string{"LtsAccessConfigInfoRespon200", string(data)}, " ")
}
