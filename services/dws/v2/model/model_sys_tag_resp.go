package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SysTagResp struct {

	// **参数解释**： 集群ID。 **取值范围**： 不涉及。
	ClusterId *string `json:"cluster_id,omitempty"`

	// **参数解释**： 企业项目ID。 **取值范围**： 不涉及。
	EpsId *string `json:"eps_id,omitempty"`
}

func (o SysTagResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SysTagResp struct{}"
	}

	return strings.Join([]string{"SysTagResp", string(data)}, " ")
}
