package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListVideoScriptsResponse Response Object
type ListVideoScriptsResponse struct {

	// 剧本总数。
	Count *int32 `json:"count,omitempty"`

	// 剧本列表。
	VideoScripts *[]VideoScriptBaseInfo `json:"video_scripts,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListVideoScriptsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVideoScriptsResponse struct{}"
	}

	return strings.Join([]string{"ListVideoScriptsResponse", string(data)}, " ")
}
