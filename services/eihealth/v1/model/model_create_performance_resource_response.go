package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePerformanceResourceResponse Response Object
type CreatePerformanceResourceResponse struct {

	// **参数解释**： 性能加速ID。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	Id             *string `json:"id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreatePerformanceResourceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePerformanceResourceResponse struct{}"
	}

	return strings.Join([]string{"CreatePerformanceResourceResponse", string(data)}, " ")
}
