package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateClusterInstanceBody 实例对象。
type CreateClusterInstanceBody struct {

	// 实例规格名称。可以使用[获取实例规格列表](ListFlavors.xml)的name属性确认当前拥有的规格信息。
	FlavorRef string `json:"flavorRef"`

	Volume *CreateClusterInstanceVolumeBody `json:"volume"`

	Nics *CreateClusterInstanceNicsBody `json:"nics"`

	AvailabilityZone *string `json:"availability_zone,omitempty"`
}

func (o CreateClusterInstanceBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateClusterInstanceBody struct{}"
	}

	return strings.Join([]string{"CreateClusterInstanceBody", string(data)}, " ")
}
