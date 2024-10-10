package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BackupFileInfo 备份文件信息。
type BackupFileInfo struct {

	// 备份文件名称。
	Name string `json:"name"`

	// OBS桶中文件路径。  OBS场景：必选  RDS场景：缺省
	ObsPath *string `json:"obs_path,omitempty"`

	// bak文件数据库版本。  OBS场景：缺省  RDS场景：必选
	RdsVersion *string `json:"rds_version,omitempty"`

	// bak文件所属实例。  OBS场景：缺省  RDS场景：必选
	RdsSourceInstanceId *string `json:"rds_source_instance_id,omitempty"`
}

func (o BackupFileInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BackupFileInfo struct{}"
	}

	return strings.Join([]string{"BackupFileInfo", string(data)}, " ")
}
