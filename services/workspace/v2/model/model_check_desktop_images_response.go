package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckDesktopImagesResponse Response Object
type CheckDesktopImagesResponse struct {

	// 桌面镜像信息。
	ImageInfos     *[]DesktopImageInfo `json:"image_infos,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o CheckDesktopImagesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckDesktopImagesResponse struct{}"
	}

	return strings.Join([]string{"CheckDesktopImagesResponse", string(data)}, " ")
}
