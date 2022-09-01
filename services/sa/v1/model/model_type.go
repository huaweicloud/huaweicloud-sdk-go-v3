package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Type struct {

	// 事件所属业务领域标签，可选类别如下： attack – 攻击 vulnerability – 漏洞 compliance check – 合规检查 risk - 风险 public opinion - 舆情 illegal&violation - 违法违规 security bulletin - 公告
	Business string `json:"business" xml:"business"`

	// 类别，推荐使用预定义的类型分类。
	Category *string `json:"category,omitempty" xml:"category"`

	// 分类器，推荐使用预定义的分类器。 如果指定了分类器，则必须指定类别。
	Classifier *string `json:"classifier,omitempty" xml:"classifier"`

	// 技术领域标签： OS：主机 APP：应用 NET：网络 OPS：运维 CS：云服务 CSP：平台云服务
	TechDomain *string `json:"tech_domain,omitempty" xml:"tech_domain"`

	Properties *TypeProperties `json:"properties,omitempty" xml:"properties"`
}

func (o Type) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Type struct{}"
	}

	return strings.Join([]string{"Type", string(data)}, " ")
}
