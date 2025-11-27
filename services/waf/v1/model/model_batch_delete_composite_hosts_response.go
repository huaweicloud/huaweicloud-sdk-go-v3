package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteCompositeHostsResponse Response Object
type BatchDeleteCompositeHostsResponse struct {

	// **参数解释：** 域名id数组，返回已成功批量删除的租户域名唯一标识。 **约束限制：** 不涉及 **取值范围：** 数组内元素为字符串类型，与请求体中传入的有效域名ID对应 **默认取值：** 不涉及
	HostIds        *[]string `json:"host_ids,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o BatchDeleteCompositeHostsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteCompositeHostsResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteCompositeHostsResponse", string(data)}, " ")
}
