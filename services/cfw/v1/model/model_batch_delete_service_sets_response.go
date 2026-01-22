package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteServiceSetsResponse Response Object
type BatchDeleteServiceSetsResponse struct {

	// **参数解释**： 服务组信息 **取值范围**： 不涉及
	Data           *[]BatchDeleteServiceSetsRespData `json:"data,omitempty"`
	HttpStatusCode int                               `json:"-"`
}

func (o BatchDeleteServiceSetsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteServiceSetsResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteServiceSetsResponse", string(data)}, " ")
}
