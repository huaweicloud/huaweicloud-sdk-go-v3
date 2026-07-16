package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TagRequest Lite Server超节点创建以及删除的标签列表。
type TagRequest struct {

	// **参数解释**：标签列表。
	Tags []TmsTag `json:"tags"`
}

func (o TagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TagRequest struct{}"
	}

	return strings.Join([]string{"TagRequest", string(data)}, " ")
}
