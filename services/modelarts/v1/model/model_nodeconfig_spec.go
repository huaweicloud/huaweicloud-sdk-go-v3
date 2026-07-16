package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type NodeconfigSpec struct {

	// **参数解释**： 节点自定义配置；当前支持节点绑核、是否启用缓存清理、是否启用透明大页。 \"configs\": {    \"cpu_manager\": {     \"mode\": \"static/none\" //static为启用绑核, none为不启用绑核    },    \"drop_cache\": {     \"mode\": \"enable/disable\" // enable启用缓存清理    },    \"transparent_hugepage\": {     \"mode\": \"always/madvise/never\" // always为启用透明大页，never为关闭透明大页，madvice交给系统选择。    } }
	Configs *interface{} `json:"configs"`
}

func (o NodeconfigSpec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodeconfigSpec struct{}"
	}

	return strings.Join([]string{"NodeconfigSpec", string(data)}, " ")
}
