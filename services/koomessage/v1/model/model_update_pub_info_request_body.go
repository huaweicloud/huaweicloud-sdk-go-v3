package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdatePubInfoRequestBody 更新服务号请求体。
type UpdatePubInfoRequestBody struct {

	// 更新原因。
	ChangeReason string `json:"change_reason"`

	// 服务号LOGO图片资源ID。  > 通过上传智能信息服务号图片资源接口上传LOGO图片后获得的图片resource_id。图片要求大于等于240*240且比例相同。
	LogoImg *string `json:"logo_img,omitempty"`

	// 授权证明图片资源ID，最多支持6张。
	AuthorizationFiles *[]string `json:"authorization_files,omitempty"`

	// 服务号名称。
	PubName *string `json:"pub_name,omitempty"`

	// 服务号简介。
	PubAbstract *string `json:"pub_abstract,omitempty"`

	// 服务号备注。
	PubRemark *string `json:"pub_remark,omitempty"`

	// 自动获取端口。
	AutoGetPort *int32 `json:"auto_get_port,omitempty"`

	// 自动收集端口使用的签名列表。  > auto_get_port为1时，该字段为必填，每个签名长度须为2-18个字符，每个服务号签名不可以重复。
	SignsForAutoGetPort *[]string `json:"signs_for_auto_get_port,omitempty"`

	// 自动收集端口生效的地区列表。地区取值见《地区名称列表》。  > auto_get_port为1时，该字段有效。不填则默认全国，不允许传入重叠地区。
	AreasForAutoGetPort *[]string `json:"areas_for_auto_get_port,omitempty"`

	// 从事行业，默认取服务号所属商家的行业分类。 - 1：金融理财 - 2：社交通讯 - 3：影音娱乐 - 4：旅游出行 - 5：购物 - 6：本地生活 - 7：运动健康 - 8：教育培训 - 9：新闻阅读 - 10：运营商  - 11：其他
	Industry *int32 `json:"industry,omitempty"`
}

func (o UpdatePubInfoRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePubInfoRequestBody struct{}"
	}

	return strings.Join([]string{"UpdatePubInfoRequestBody", string(data)}, " ")
}
