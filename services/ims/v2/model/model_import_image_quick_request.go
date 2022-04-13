package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ImportImageQuickRequest struct {
	Body *QuickImportImageByFileRequestBody `json:"body,omitempty"`
}

func (o ImportImageQuickRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportImageQuickRequest struct{}"
	}

	return strings.Join([]string{"ImportImageQuickRequest", string(data)}, " ")
}
