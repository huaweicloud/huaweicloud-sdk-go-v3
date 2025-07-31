package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImportImageQuickInSafeModeRequest Request Object
type ImportImageQuickInSafeModeRequest struct {
	Body *QuickImportImageByFileRequestBody `json:"body,omitempty"`
}

func (o ImportImageQuickInSafeModeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportImageQuickInSafeModeRequest struct{}"
	}

	return strings.Join([]string{"ImportImageQuickInSafeModeRequest", string(data)}, " ")
}
