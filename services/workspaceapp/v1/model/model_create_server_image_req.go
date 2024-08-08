package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateServerImageReq 构建云应用镜像的请求体。
type CreateServerImageReq struct {

	// 镜像名称，名称需满足如下规则: * 首尾字符不能为空格。 * 长度范围1～128个字符。 * 只允许3种字符，英文大小写，数字，特殊字符包含-、.、_、空格和中文。
	Name string `json:"name"`

	// 镜像描述，描述需满足如下规则: * 首字符不能为空格。 * 长度范围0～1024个字符。 * 支持字母、数字、中文。 * 不支持回车、<、 >字符。
	Description *string `json:"description,omitempty"`

	// **⚠ : 此属性是预留字段，不需要传值，目前镜像产物默认属于default企业项目** 镜像所属的企业项目ID，默认属于default企业项目。 关于企业项目ID的获取及企业项目特性的详细信息，请参考“[企业中心总览](https://support.huaweicloud.com/intl/zh-cn/usermanual-em/zh-cn_topic_0123692049.html)”。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o CreateServerImageReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateServerImageReq struct{}"
	}

	return strings.Join([]string{"CreateServerImageReq", string(data)}, " ")
}
