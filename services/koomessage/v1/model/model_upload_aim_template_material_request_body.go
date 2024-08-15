package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UploadAimTemplateMaterialRequestBody struct {

	// 资源类型。  - image：图片 - video：视频 - thumbnail：缩略图  > 图片支持png、jpeg、jpg格式，最大2M; > 视频支持格式mp4，大小不超过7M，建议时长不超过33S; > 缩略图支持png、jpeg、jpg格式，大小不超过100K。
	ResourceType string `json:"resource_type"`

	// 素材ID。 > resource_type=thumbnail时，material_id必填，填写内容为上传视频资源返回的material_id字段的值。
	MaterialId *string `json:"material_id,omitempty"`

	// 文件类型。 - url：资源为URL链接地址 - stream：资源为多媒体资源文件流的BASE64编码，需要带资源类型前缀，样例：\"data:image/jpeg;base64,/9j/4AAQSkZJRgABAQAAAQABAAD/4gIoSUNDX1BST0ZJTEUAAQEAAAIYAAAAAAQwAABtbnRyUkdCIFhZWiAAAAAAAAAAAAAAAABhY3NwAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAQAA9tYAAQAAAADTLQ...\"，样例过长，未显示全部
	FileType string `json:"file_type"`

	// 文件URL。  > file_type=url时，file_url为必填。URL地址必须包含文件格式的后缀，例如：jpg、jpeg，大小写后缀都支持。
	FileUrl *string `json:"file_url,omitempty"`

	// 文件资源码流。  > file_type=stream时，file_stream为必填。
	FileStream *string `json:"file_stream,omitempty"`

	// 图像比例。 - oneToOne：指1:1比例 - sixteenToNine：指16:9比例 - threeToOne：指3:1比例 - fortyEightToSixtyFive：指48:65比例 - twentyOneToNine：指21:9比例 - threeToFour：指3:4比例  > resource type=image时，image_rate必填。
	ImageRate *string `json:"image_rate,omitempty"`

	// 文件名称。  > file_type=stream时，file_name必填。
	FileName *string `json:"file_name,omitempty"`

	// 描述。
	Description *string `json:"description,omitempty"`
}

func (o UploadAimTemplateMaterialRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UploadAimTemplateMaterialRequestBody struct{}"
	}

	return strings.Join([]string{"UploadAimTemplateMaterialRequestBody", string(data)}, " ")
}
