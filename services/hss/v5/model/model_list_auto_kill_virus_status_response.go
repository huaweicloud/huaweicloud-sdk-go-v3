package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAutoKillVirusStatusResponse Response Object
type ListAutoKillVirusStatusResponse struct {

	// **参数解释**： 程序自动隔离查杀开关状态 **取值范围**： - true：开启 - false：关闭
	Enabled        *bool `json:"enabled,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o ListAutoKillVirusStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAutoKillVirusStatusResponse struct{}"
	}

	return strings.Join([]string{"ListAutoKillVirusStatusResponse", string(data)}, " ")
}
