package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 配额信息。
type QuotaList struct {
	BackupGigabytes *QuotaDetailBackupGigabytes `json:"backup_gigabytes"`

	Backups *QuotaDetailBackups `json:"backups"`

	Gigabytes *QuotaDetailGigabytes `json:"gigabytes"`
	// 项目ID。

	Id string `json:"id"`

	Snapshots *QuotaDetailSnapshots `json:"snapshots"`

	Volumes *QuotaDetailVolumes `json:"volumes"`

	GigabytesSATA *QuotaDetailGigabytesSata `json:"gigabytes_SATA,omitempty"`

	SnapshotsSATA *QuotaDetailSnapshotsSata `json:"snapshots_SATA,omitempty"`

	VolumesSATA *QuotaDetailVolumesSata `json:"volumes_SATA,omitempty"`

	GigabytesSAS *QuotaDetailGigabytesSas `json:"gigabytes_SAS,omitempty"`

	SnapshotsSAS *QuotaDetailSnapshotsSas `json:"snapshots_SAS,omitempty"`

	VolumesSAS *QuotaDetailVolumesSas `json:"volumes_SAS,omitempty"`

	GigabytesSSD *QuotaDetailGigabytesSsd `json:"gigabytes_SSD,omitempty"`

	SnapshotsSSD *QuotaDetailSnapshotsSsd `json:"snapshots_SSD,omitempty"`

	VolumesSSD *QuotaDetailVolumesSsd `json:"volumes_SSD,omitempty"`

	GigabytesGPSSD *QuotaDetailGigabytesGpssd `json:"gigabytes_GPSSD,omitempty"`

	SnapshotsGPSSD *QuotaDetailSnapshotsGpssd `json:"snapshots_GPSSD,omitempty"`

	VolumesGPSSD *QuotaDetailVolumesGpssd `json:"volumes_GPSSD,omitempty"`

	PerVolumeGigabytes *QuotaDetailPerVolumeGigabytes `json:"per_volume_gigabytes,omitempty"`
}

func (o QuotaList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QuotaList struct{}"
	}

	return strings.Join([]string{"QuotaList", string(data)}, " ")
}
