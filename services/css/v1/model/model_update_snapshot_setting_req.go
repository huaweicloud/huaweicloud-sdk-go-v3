package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateSnapshotSettingReq struct {
	// 备份使用的OBS桶，如果桶已经存放快照数据了，不可更改。

	Bucket string `json:"bucket"`
	// 访问OBS使用的IAM委托名称。

	Agency string `json:"agency"`
	// 快照在OBS桶中的存放路径。

	BasePath string `json:"basePath"`
}

func (o UpdateSnapshotSettingReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSnapshotSettingReq struct{}"
	}

	return strings.Join([]string{"UpdateSnapshotSettingReq", string(data)}, " ")
}
