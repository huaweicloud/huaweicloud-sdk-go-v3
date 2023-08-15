package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DetectLiveByBase64Request Request Object
type DetectLiveByBase64Request struct {

	// 企业项目ID。FRS支持通过企业项目管理（EPS）对不同用户组和用户的资源使用，进行分账。当前仅支持按需计费模式。  获取方法：进入“[企业项目管理](https://console.huaweicloud.com/eps/?region=cn-north-4#/projects/list)”页面，单击企业项目名称，在企业项目详情页获取Enterprise-Project-Id（企业项目ID）。  企业项目创建步骤请参见用户指南。 > 说明： 创建企业项目后，在传参时，有以下三类场景。 - 携带正确的ID，正常使用FRS服务，账单归到企业ID对应的企业项目中。 - 携带错误的ID，正常使用FRS服务，账单的企业项目会被分类为“未归集”。 - 不携带ID，正常使用FRS服务，账单的企业项目会被分类为“未归集”。
	EnterpriseProjectId *string `json:"Enterprise-Project-Id,omitempty"`

	Body *LiveDetectBase64Req `json:"body,omitempty"`
}

func (o DetectLiveByBase64Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DetectLiveByBase64Request struct{}"
	}

	return strings.Join([]string{"DetectLiveByBase64Request", string(data)}, " ")
}
