package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteDomainSetResponse Response Object
type BatchDeleteDomainSetResponse struct {

	// **参数解释**： 批量删除域名组响应信息 **取值范围**： 不涉及
	Data           map[string][]Mapstringstring `json:"data,omitempty"`
	HttpStatusCode int                          `json:"-"`
}

func (o BatchDeleteDomainSetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteDomainSetResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteDomainSetResponse", string(data)}, " ")
}
