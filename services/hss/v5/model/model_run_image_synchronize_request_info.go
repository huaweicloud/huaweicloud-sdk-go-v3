package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RunImageSynchronizeRequestInfo struct {

	// **参数解释**: 镜像类型 **约束限制**: 必选参数，仅支持指定取值 **取值范围**: - private_image : 私有镜像仓库 - shared_image : 共享镜像仓库 **默认取值**: 不涉及
	ImageType string `json:"image_type"`
}

func (o RunImageSynchronizeRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunImageSynchronizeRequestInfo struct{}"
	}

	return strings.Join([]string{"RunImageSynchronizeRequestInfo", string(data)}, " ")
}
