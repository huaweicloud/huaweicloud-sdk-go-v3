package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetAutoEnlargePolicyResponse Response Object
type SetAutoEnlargePolicyResponse struct {

	// **参数解释：** 设置磁盘自动扩容策略失败的实例信息列表。 **取值范围：** 不涉及
	ErrorResults   *[]SetAutoPolicyErrorResults `json:"error_results,omitempty"`
	HttpStatusCode int                          `json:"-"`
}

func (o SetAutoEnlargePolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetAutoEnlargePolicyResponse struct{}"
	}

	return strings.Join([]string{"SetAutoEnlargePolicyResponse", string(data)}, " ")
}
