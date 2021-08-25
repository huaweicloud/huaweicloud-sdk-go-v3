package model

import (
	"encoding/json"

	"strings"
)

// 添加图片body体
type AddPictureRequestReq struct {
	// 图片文件Base64编码字符串，仅支持JPEG/JPG/PNG/BMP格式，图片最小边不小于100px，最大边不超过2048px。

	File *string `json:"file,omitempty"`
	// 图片的URL路径，作为图片库中索引图片的ID，是必选参数。  > - 当file字段不为空时，图片从file获取，path作为图片索引ID使用；当file字段不存在或者为空时，图片需要通过下载获取，此时path作为下载图片的地址（当前仅支持从华为云图像搜索服务所在区域的OBS下载图片），同时，path也作为图片索引ID。

	Path string `json:"path"`
	// 图片自定义标签。格式为key：value对，所有图片的key总数最多不超过10个，但是每个key对应的value不限制个数，例如：key为动物，对应的value可以是猫、狗、鸟等多个。  标签名（key）添加方式：   - 登录管理控制台，单击“创建实例”，自定义标签名。    - 登录管理控制台，在“实例管理”页签，单击实例名称，进入“基础信息”页添加自定义标签。

	Tags *interface{} `json:"tags,omitempty"`
}

func (o AddPictureRequestReq) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "AddPictureRequestReq struct{}"
	}

	return strings.Join([]string{"AddPictureRequestReq", string(data)}, " ")
}
