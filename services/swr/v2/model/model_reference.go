package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Reference struct {

	// 父制品的ID。
	ParentId *int64 `json:"parent_id,omitempty"`

	// 子制品的ID。
	ChildId *int64 `json:"child_id,omitempty"`

	// 子制品的digest。
	ChildDigest *string `json:"child_digest,omitempty"`

	Platform *Platform `json:"platform,omitempty"`
}

func (o Reference) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Reference struct{}"
	}

	return strings.Join([]string{"Reference", string(data)}, " ")
}
