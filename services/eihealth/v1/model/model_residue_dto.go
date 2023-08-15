package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResidueDto 氨基酸残基或者配体
type ResidueDto struct {

	// 氨基酸残基或者配体链的名称
	Chain string `json:"chain"`

	// 氨基酸残基或者配体名称
	Name *string `json:"name,omitempty"`

	// 氨基酸残基或者配体的序列ID
	Id *int64 `json:"id,omitempty"`
}

func (o ResidueDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResidueDto struct{}"
	}

	return strings.Join([]string{"ResidueDto", string(data)}, " ")
}
