package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddTagToAssetRequest Request Object
type AddTagToAssetRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	// guid
	TermGuid string `json:"term_guid"`

	Body *CatalogInfo `json:"body,omitempty"`
}

func (o AddTagToAssetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddTagToAssetRequest struct{}"
	}

	return strings.Join([]string{"AddTagToAssetRequest", string(data)}, " ")
}
