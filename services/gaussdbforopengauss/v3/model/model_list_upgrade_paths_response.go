package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListUpgradePathsResponse Response Object
type ListUpgradePathsResponse struct {

	// **参数解释**: 版本信息详情
	VersionInfos *[]VersionInfosResult `json:"version_infos,omitempty"`

	// **参数解释**: 支持的升级路径列表
	VersionEdges   *[]UpgradePathsResult `json:"version_edges,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o ListUpgradePathsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListUpgradePathsResponse struct{}"
	}

	return strings.Join([]string{"ListUpgradePathsResponse", string(data)}, " ")
}
