package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Shards struct {

	// **参数解释**：日志分片的obs下载链接集。 **取值范围**：不涉及。
	ObjectUrls *[]string `json:"object_urls,omitempty"`

	// **参数解释**：本次请求返回的日志分片数。 **取值范围**：不涉及。
	Count *int32 `json:"count,omitempty"`

	// **参数解释**：本次请求命中的日志分片总数。 **取值范围**：不涉及。
	Total *int32 `json:"total,omitempty"`
}

func (o Shards) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Shards struct{}"
	}

	return strings.Join([]string{"Shards", string(data)}, " ")
}
