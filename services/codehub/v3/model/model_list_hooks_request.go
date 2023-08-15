package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListHooksRequest Request Object
type ListHooksRequest struct {

	// 组名(克隆地址中域名后面仓库名前的一段 示例：git@repo.alpha.devcloud.inhuawei.com:Demo00228/testword.git  组名：Demo00228 )
	GroupName string `json:"group_name"`

	// hook id
	HookId *string `json:"hook_id,omitempty"`

	// 仓库名
	RepositoryName string `json:"repository_name"`
}

func (o ListHooksRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHooksRequest struct{}"
	}

	return strings.Join([]string{"ListHooksRequest", string(data)}, " ")
}
