package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ProductResponse 产品详情返回体
type ProductResponse struct {

	// 产品id
	Id *string `json:"id,omitempty"`

	// 产品名称
	Name string `json:"name"`

	// 产品描述
	Description *string `json:"description,omitempty"`

	// 产品属性值
	Attributes map[string]ProductAttributes `json:"attributes,omitempty"`

	// 产品所属账号的项目ID
	ProjectId string `json:"project_id"`

	// 产品创建时间戳
	CreatedAt *int32 `json:"created_at,omitempty"`

	// 产品标签
	Tags *[]Attributes `json:"tags,omitempty"`

	// 产品私钥
	PrivateKey *string `json:"private_key,omitempty"`

	// 产品证书
	Certificate *string `json:"certificate,omitempty"`

	// 产品根证书
	Ca *string `json:"ca,omitempty"`

	// 将产品证书文件certificate/ca/private_key打成.tar.gz包后用base64编码的字符串。 使用时请使用base64解码成.tar.gz包。
	Package *string `json:"package,omitempty"`

	// 产品使用token注册时的凭证
	Identifier *string `json:"identifier,omitempty"`
}

func (o ProductResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProductResponse struct{}"
	}

	return strings.Join([]string{"ProductResponse", string(data)}, " ")
}
