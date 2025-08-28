package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImageWhiteListDetailResponseInfoQueryInfo 白名单规则为“specific_image_types”时，指定的镜像类型信息
type ImageWhiteListDetailResponseInfoQueryInfo struct {

	// **参数解释**： 镜像类型（多个类型用逗号间隔） **取值范围**： - private_image：SWR私有镜像仓库。 - shared_image：SWR共享镜像仓库。 - instance_image：SWR企业仓库。 - harbor：harbor仓库。 - jfrog：jfrog仓库。
	ImageType *string `json:"image_type,omitempty"`
}

func (o ImageWhiteListDetailResponseInfoQueryInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImageWhiteListDetailResponseInfoQueryInfo struct{}"
	}

	return strings.Join([]string{"ImageWhiteListDetailResponseInfoQueryInfo", string(data)}, " ")
}
