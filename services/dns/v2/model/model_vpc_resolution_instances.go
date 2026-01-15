package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VpcResolutionInstances struct {

	// **参数解释：** VPC ID。 **取值范围：** 不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释：** VPC名称。 **取值范围：** 不涉及。
	Name *string `json:"name,omitempty"`
}

func (o VpcResolutionInstances) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VpcResolutionInstances struct{}"
	}

	return strings.Join([]string{"VpcResolutionInstances", string(data)}, " ")
}
