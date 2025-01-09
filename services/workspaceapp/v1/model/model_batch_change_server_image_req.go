package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchChangeServerImageReq 批量修改服务器镜像的请求体。
type BatchChangeServerImageReq struct {

	// 应用服务器id集合。
	ServerIds *[]string `json:"server_ids,omitempty"`

	// 镜像id，要求与服务器原有镜像id不相同。
	ImageId string `json:"image_id"`

	ImageType *ImageTypeEnum `json:"image_type"`

	OsType *OsTypeEnum `json:"os_type"`

	// 镜像的产品id，当镜像是市场镜像时候，该字段必传。
	ImageProductId *string `json:"image_product_id,omitempty"`

	// 是否自动升级hda版本。
	UpdateAccessAgent *bool `json:"update_access_agent,omitempty"`
}

func (o BatchChangeServerImageReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchChangeServerImageReq struct{}"
	}

	return strings.Join([]string{"BatchChangeServerImageReq", string(data)}, " ")
}
