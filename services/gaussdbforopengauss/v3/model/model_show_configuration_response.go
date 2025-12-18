package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowConfigurationResponse Response Object
type ShowConfigurationResponse struct {

	// **参数解释**: 备份限速。 **取值范围**: 0-1024
	RateLimit *int32 `json:"rate_limit,omitempty"`

	// **参数解释**: 增备预取页面个数。 **取值范围**: 1-8192
	PrefetchBlock *int32 `json:"prefetch_block,omitempty"`

	// **参数解释**: 分片大小。 **取值范围**: 0-1024
	FileSplitSize *int32 `json:"file_split_size,omitempty"`

	// **参数解释**: 启用备机备份标识。 **取值范围**: 不涉及。
	EnableStandbyBackup *bool `json:"enable_standby_backup,omitempty"`

	// **参数解释**: 是否关闭压缩。 **取值范围**: true,false
	CloseCompression *bool `json:"close_compression,omitempty"`

	// **参数解释**: 默认备份介质。 **取值范围**: - OBS 对象存储
	DefaultBackupMediaType *string `json:"default_backup_media_type,omitempty"`

	// **参数解释**: 默认备份方式。 **取值范围**: - EBACKUP 快照备份 - PHYSICAL_BACKUP 物理备份
	DefaultBackupMethod *string `json:"default_backup_method,omitempty"`

	// **参数解释**: 备份并行参数。 **取值范围**: 1, 2, 4, 8
	BackupParallelDegree *int32 `json:"backup_parallel_degree,omitempty"`

	BackupNodeInfo *BackupNodeInfoResult `json:"backup_node_info,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o ShowConfigurationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowConfigurationResponse struct{}"
	}

	return strings.Join([]string{"ShowConfigurationResponse", string(data)}, " ")
}
