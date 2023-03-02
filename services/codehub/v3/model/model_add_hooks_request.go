package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type AddHooksRequest struct {

	// 组名(克隆地址中域名后面项目名前的一段 示例：git@repo.alpha.devcloud.inhuawei.com:Demo00228/testword.git  组名：Demo00228 )
	GroupName string `json:"group_name"`

	// 仓库名
	RepositoryName string `json:"repository_name"`

	Body *RepositoryHookRequest `json:"body,omitempty"`
}

func (o AddHooksRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddHooksRequest struct{}"
	}

	return strings.Join([]string{"AddHooksRequest", string(data)}, " ")
}
