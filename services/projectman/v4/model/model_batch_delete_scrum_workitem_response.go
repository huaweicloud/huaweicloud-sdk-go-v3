package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteScrumWorkitemResponse Response Object
type BatchDeleteScrumWorkitemResponse struct {
	Result *BatchDeletesResponseResult `json:"result,omitempty"`

	// **参数解释：** 返回状态。 **取值范围：** success：删除成功。 error：删除失败。
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o BatchDeleteScrumWorkitemResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteScrumWorkitemResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteScrumWorkitemResponse", string(data)}, " ")
}
