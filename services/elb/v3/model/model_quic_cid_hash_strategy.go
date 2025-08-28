package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type QuicCidHashStrategy struct {

	// **参数解释**：表示hash的时候取CID的长度。  **约束限制**：仅当负载均衡算法为QUIC_CID的时候才生效。  **取值范围**：1-20  **默认取值**：3
	Len int32 `json:"len"`

	// **参数解释**：表示hash的时候取CID的偏移量。  **约束限制**：仅当负载均衡算法为QUIC_CID的时候才生效。  **取值范围**：0-19  **默认取值**：1
	Offset int32 `json:"offset"`
}

func (o QuicCidHashStrategy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QuicCidHashStrategy struct{}"
	}

	return strings.Join([]string{"QuicCidHashStrategy", string(data)}, " ")
}
