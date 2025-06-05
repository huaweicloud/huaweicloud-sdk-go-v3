package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListClusterInstallStepResponse Response Object
type ListClusterInstallStepResponse struct {

	// 安装步骤详情列表。
	Installs *[]InstallStep `json:"installs,omitempty"`

	// 卸载步骤详情列表。
	Uninstalls     *[]InstallStep `json:"uninstalls,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ListClusterInstallStepResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListClusterInstallStepResponse struct{}"
	}

	return strings.Join([]string{"ListClusterInstallStepResponse", string(data)}, " ")
}
