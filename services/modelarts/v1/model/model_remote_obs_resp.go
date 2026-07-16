package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RemoteObsResp 数据实际输出到OBS。
type RemoteObsResp struct {

	// **参数解释**：数据实际输出到OBS的路径。 **取值范围**：不涉及。
	ObsUrl string `json:"obs_url"`
}

func (o RemoteObsResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RemoteObsResp struct{}"
	}

	return strings.Join([]string{"RemoteObsResp", string(data)}, " ")
}
