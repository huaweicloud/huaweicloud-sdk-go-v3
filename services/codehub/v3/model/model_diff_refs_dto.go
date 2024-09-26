package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DiffRefsDto struct {

	// base sha
	BaseSha *string `json:"base_sha,omitempty"`

	// head sha
	HeadSha *string `json:"head_sha,omitempty"`

	// start sha
	StartSha *string `json:"start_sha,omitempty"`
}

func (o DiffRefsDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DiffRefsDto struct{}"
	}

	return strings.Join([]string{"DiffRefsDto", string(data)}, " ")
}
