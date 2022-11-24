package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ExportCertificateRequestBody struct {

	// 是否压缩。   - **true**   - **false**
	IsCompressed string `json:"is_compressed"`

	// 根据服务器类型选择下载证书的形式，支持以下五种类型：   - **APACHE** : apache服务器推荐使用此参数；   - **NGINX** : nginx服务器推荐使用此参数；   - **IIS** : windows服务器推荐使用此参数；   - **TOMCAT** : tomcat服务器推荐使用此参数；   - **OTHER** : 下载PEM格式证书，推荐使用此参数。
	Type string `json:"type"`
}

func (o ExportCertificateRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportCertificateRequestBody struct{}"
	}

	return strings.Join([]string{"ExportCertificateRequestBody", string(data)}, " ")
}
