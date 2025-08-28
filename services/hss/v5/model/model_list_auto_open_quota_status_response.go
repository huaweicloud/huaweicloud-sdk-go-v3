package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAutoOpenQuotaStatusResponse Response Object
type ListAutoOpenQuotaStatusResponse struct {

	// **参数解释**： 自动绑定配额配置开关状态 **取值范围**： - true：开启 - false：关闭
	Enabled        *bool `json:"enabled,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o ListAutoOpenQuotaStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAutoOpenQuotaStatusResponse struct{}"
	}

	return strings.Join([]string{"ListAutoOpenQuotaStatusResponse", string(data)}, " ")
}
