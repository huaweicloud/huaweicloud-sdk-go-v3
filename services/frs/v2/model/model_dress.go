package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Dress struct {
	// 是否带眼镜： • yes：带眼镜 • dark：带墨镜 • none：未戴眼镜 • unknown：未知

	Glass *string `json:"glass,omitempty"`
	// 是否戴帽子： • yes：戴帽子 • none：未戴帽子 • unknown：未知

	Hat *string `json:"hat,omitempty"`
}

func (o Dress) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Dress struct{}"
	}

	return strings.Join([]string{"Dress", string(data)}, " ")
}
