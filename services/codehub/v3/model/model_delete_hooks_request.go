package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteHooksRequest Request Object
type DeleteHooksRequest struct {

	// 组名(克隆地址中域名后面项目名前的一段 示例：git@repo.alpha.devcloud.inhuawei.com:Demo00228/testword.git  组名：Demo00228 )
	GroupName string `json:"group_name"`

	// 通过id删除指定仓库的hook
	HookId int32 `json:"hook_id"`

	// 仓库名
	RepositoryName string `json:"repository_name"`
}

func (o DeleteHooksRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteHooksRequest struct{}"
	}

	return strings.Join([]string{"DeleteHooksRequest", string(data)}, " ")
}
