package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListFactoryScriptsResponse Response Object
type ListFactoryScriptsResponse struct {

	// 总的脚本个数
	Total *int32 `json:"total,omitempty"`

	// 脚本列表
	Scripts        *[]ScriptInfo `json:"scripts,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ListFactoryScriptsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFactoryScriptsResponse struct{}"
	}

	return strings.Join([]string{"ListFactoryScriptsResponse", string(data)}, " ")
}
