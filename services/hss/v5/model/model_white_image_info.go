package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WhiteImageInfo 白名单镜像
type WhiteImageInfo struct {

	// **参数解释**： 集群ID **取值范围**： 不涉及
	ClusterId string `json:"cluster_id"`

	// **参数解释**： 镜像名称 **取值范围**： 不涉及
	ImageName string `json:"image_name"`

	// **参数解释**： 镜像版本 **取值范围**： 不涉及
	ImageVersion string `json:"image_version"`
}

func (o WhiteImageInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WhiteImageInfo struct{}"
	}

	return strings.Join([]string{"WhiteImageInfo", string(data)}, " ")
}
