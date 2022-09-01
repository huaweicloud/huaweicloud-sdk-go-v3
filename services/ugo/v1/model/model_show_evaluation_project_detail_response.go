package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowEvaluationProjectDetailResponse struct {

	// 数据库大小。
	SourceDbSize *string `json:"source_db_size,omitempty" xml:"source_db_size"`

	// 数据库schema个数。
	SourceDbSchema *int64 `json:"source_db_schema,omitempty" xml:"source_db_schema"`

	// 数据库CPU个数。
	SourceDbCpu *string `json:"source_db_cpu,omitempty" xml:"source_db_cpu"`

	// 数据库字符集。
	SourceDbCharacterSet *string `json:"source_db_character_set,omitempty" xml:"source_db_character_set"`

	// 数据库操作系统。
	SourceDbOs *string `json:"source_db_os,omitempty" xml:"source_db_os"`

	// 实例数量。
	SourceDbInstanceNum *int32 `json:"source_db_instance_num,omitempty" xml:"source_db_instance_num"`

	// 数据库内存。
	SourceDbRam *string `json:"source_db_ram,omitempty" xml:"source_db_ram"`

	SourceDbInfo *SourceDb `json:"source_db_info,omitempty" xml:"source_db_info"`

	// 数据库物理RAM。
	SourceDbPhyRam *string `json:"source_db_phy_ram,omitempty" xml:"source_db_phy_ram"`

	// 数据库版本。
	SourceDbVersion *string `json:"source_db_version,omitempty" xml:"source_db_version"`

	// 数据库配置。
	SourceDbConf *string `json:"source_db_conf,omitempty" xml:"source_db_conf"`

	// 数据库时区。
	SourceDbClock *string `json:"source_db_clock,omitempty" xml:"source_db_clock"`

	// 评估项目ID。
	EvaluationProjectId *int32 `json:"evaluation_project_id,omitempty" xml:"evaluation_project_id"`

	// 评估项目名称。
	EvaluationProjectName *string `json:"evaluation_project_name,omitempty" xml:"evaluation_project_name"`
	HttpStatusCode        int     `json:"-"`
}

func (o ShowEvaluationProjectDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowEvaluationProjectDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowEvaluationProjectDetailResponse", string(data)}, " ")
}
