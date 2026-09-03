package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PackageRegionWithStatusInfo 技能包区域信息（企业自研技能，含上传状态）。服务端根据 slug+version+packageName+region 自动构造 OBS 路径。
type PackageRegionWithStatusInfo struct {

	// 区域标识（如 cn-north-7）。
	Region string `json:"region"`

	UploadStatus *UploadStatusEnum `json:"upload_status"`
}

func (o PackageRegionWithStatusInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PackageRegionWithStatusInfo struct{}"
	}

	return strings.Join([]string{"PackageRegionWithStatusInfo", string(data)}, " ")
}
