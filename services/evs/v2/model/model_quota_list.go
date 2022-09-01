package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 配额信息。
type QuotaList struct {
	BackupGigabytes *QuotaDetailBackupGigabytes `json:"backup_gigabytes" xml:"backup_gigabytes"`

	Backups *QuotaDetailBackups `json:"backups" xml:"backups"`

	Gigabytes *QuotaDetailGigabytes `json:"gigabytes" xml:"gigabytes"`

	// 项目ID。
	Id string `json:"id" xml:"id"`

	Snapshots *QuotaDetailSnapshots `json:"snapshots" xml:"snapshots"`

	Volumes *QuotaDetailVolumes `json:"volumes" xml:"volumes"`

	GigabytesSATA *QuotaDetailGigabytesSata `json:"gigabytes_SATA,omitempty" xml:"gigabytes_SATA"`

	SnapshotsSATA *QuotaDetailSnapshotsSata `json:"snapshots_SATA,omitempty" xml:"snapshots_SATA"`

	VolumesSATA *QuotaDetailVolumesSata `json:"volumes_SATA,omitempty" xml:"volumes_SATA"`

	GigabytesSAS *QuotaDetailGigabytesSas `json:"gigabytes_SAS,omitempty" xml:"gigabytes_SAS"`

	SnapshotsSAS *QuotaDetailSnapshotsSas `json:"snapshots_SAS,omitempty" xml:"snapshots_SAS"`

	VolumesSAS *QuotaDetailVolumesSas `json:"volumes_SAS,omitempty" xml:"volumes_SAS"`

	GigabytesSSD *QuotaDetailGigabytesSsd `json:"gigabytes_SSD,omitempty" xml:"gigabytes_SSD"`

	SnapshotsSSD *QuotaDetailSnapshotsSsd `json:"snapshots_SSD,omitempty" xml:"snapshots_SSD"`

	VolumesSSD *QuotaDetailVolumesSsd `json:"volumes_SSD,omitempty" xml:"volumes_SSD"`

	GigabytesGPSSD *QuotaDetailGigabytesGpssd `json:"gigabytes_GPSSD,omitempty" xml:"gigabytes_GPSSD"`

	SnapshotsGPSSD *QuotaDetailSnapshotsGpssd `json:"snapshots_GPSSD,omitempty" xml:"snapshots_GPSSD"`

	VolumesGPSSD *QuotaDetailVolumesGpssd `json:"volumes_GPSSD,omitempty" xml:"volumes_GPSSD"`

	PerVolumeGigabytes *QuotaDetailPerVolumeGigabytes `json:"per_volume_gigabytes,omitempty" xml:"per_volume_gigabytes"`
}

func (o QuotaList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QuotaList struct{}"
	}

	return strings.Join([]string{"QuotaList", string(data)}, " ")
}
