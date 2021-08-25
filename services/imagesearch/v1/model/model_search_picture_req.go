package model

import (
	"encoding/json"

	"strings"
)

type SearchPictureReq struct {
	// 图片文件Base64编码字符串，仅支持JPEG/JPG/PNG/BMP格式，图片最小边不小于100px，最大边不超过2048px。该参数与path二选一，如果两个参数都存在，则以file为主。

	File *string `json:"file,omitempty"`
	// 图片的URL路径，图片库中的图片索引ID。该参数与file二选一，如果两个参数都存在，则以file为主。

	Path *string `json:"path,omitempty"`
	// 返回被检索图像的数量，取值为1~100的整数，默认为10。

	Limit *int32 `json:"limit,omitempty"`
	// 偏移量，指定搜索结果返回起始位置，取值为大于或等于0的整数，默认为0。

	Offset *int32 `json:"offset,omitempty"`
	// 图片自定义标签，最多不超过10个，格式为key：value对。 标签名（key）添加方式：   - 登录管理控制台，单击“创建实例”，自定义标签名。   - 登录管理控制台，在“实例管理”页签，单击实例名称，进入“基础信息”页添加自定义标签。 使用图片标签搜索时该参数必选。

	Tags *interface{} `json:"tags,omitempty"`
	// 是否用图片中指定区域（参数box）进行搜索。默认为false，该参数目前仅对某些特定模型有效，其他模型暂不支持目标检测。 - true：用图片中指定区域（参数box）进行搜索。 - false：用全图进行搜索。

	IsCrop *bool `json:"isCrop,omitempty"`
	// 区域中x坐标的最小值，单位：像素。

	X *int32 `json:"x,omitempty"`
	// 区域中y坐标的最小值，单位：像素。

	Y *int32 `json:"y,omitempty"`
	// 区域的宽度，单位：像素。

	Width *int32 `json:"width,omitempty"`
	// 区域的高度，单位：像素。

	Height *int32 `json:"height,omitempty"`
}

func (o SearchPictureReq) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "SearchPictureReq struct{}"
	}

	return strings.Join([]string{"SearchPictureReq", string(data)}, " ")
}
