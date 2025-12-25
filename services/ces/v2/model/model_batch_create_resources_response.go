package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateResourcesResponse Response Object
type BatchCreateResourcesResponse struct {

	// **参数解释**： 成功添加的资源数目。     **取值范围**： 取值范围为0-1000。
	SucceedCount   *int32 `json:"succeed_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o BatchCreateResourcesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateResourcesResponse struct{}"
	}

	return strings.Join([]string{"BatchCreateResourcesResponse", string(data)}, " ")
}
