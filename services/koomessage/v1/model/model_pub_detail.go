package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PubDetail 服务号详情。
type PubDetail struct {

	// 服务号ID。
	PubId *string `json:"pub_id,omitempty"`

	// 最新操作时间，格式：yyyy-MM-ddTHH:mm:ssZ。
	OperTime *string `json:"oper_time,omitempty"`

	// 服务号状态。  - 1：未生效 - 2：已生效 - 3：已失效 - 4：已冻结
	State *int32 `json:"state,omitempty"`

	// 上线时间，格式为：yyyy-MM-ddTHH:mm:ssZ。
	OnlineTime *string `json:"online_time,omitempty"`

	// 企业名称。
	CompanyName *string `json:"company_name,omitempty"`

	// 服务号名称。
	PubName *string `json:"pub_name,omitempty"`

	// 服务号LOGO图片资源ID。
	LogoImg *string `json:"logo_img,omitempty"`

	// 服务号LOGO图片URL。
	LogoUrl *string `json:"logo_url,omitempty"`

	// 授权证明图片的OBSURL地址。
	AuthorizationFiles map[string]string `json:"authorization_files,omitempty"`

	// 是否授权系统自动收集端口。   - 0：否  - 1：是
	AutoGetPort *int32 `json:"auto_get_port,omitempty"`

	// 从事行业，默认取服务号所属商家的行业分类。  - 1：金融理财  - 2：社交通讯  - 3：影音娱乐  - 4：旅游出行  - 5：购物  - 6：本地生活  - 7：运动健康  - 8：教育培训  - 9：新闻阅读  - 10：运营商  - 11：其他
	Industry *int32 `json:"industry,omitempty"`

	// 服务号简介。
	PubAbstract *string `json:"pub_abstract,omitempty"`

	// 自动收集端口使用的签名列表。  > auto_get_port为1时，该字段为必填，每个签名长度为2-18个字符，每个服务号签名不可以重复。
	SignsForAutoGetPort *[]string `json:"signs_for_auto_get_port,omitempty"`

	// 企业ID。
	CompanyId *string `json:"company_id,omitempty"`

	// 服务号备注。
	PubRemark *string `json:"pub_remark,omitempty"`

	// 审核状态。 - 1：审核中 - 2：审核通过 - 3：驳回
	ApproveState *int32 `json:"approve_state,omitempty"`

	// 自动收集端口生效的地区列表。地区取值见《地区名称列表》。
	AreasForAutoGetPort *[]string `json:"areas_for_auto_get_port,omitempty"`
}

func (o PubDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PubDetail struct{}"
	}

	return strings.Join([]string{"PubDetail", string(data)}, " ")
}
