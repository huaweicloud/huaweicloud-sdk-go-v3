package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ContactWayInfoV2 struct {

	// 姓名或名称
	Name string `json:"name"`

	// 联系电话
	Mobile string `json:"mobile"`

	// 联系邮箱
	Mail *string `json:"mail,omitempty"`

	// 国家/地区
	Area *string `json:"area,omitempty"`

	// 地址
	Address *string `json:"address,omitempty"`

	// 身份。企业 10；个人 20；授权代理人 21；律师 22；知识产权所有人 23
	IdType int32 `json:"id_type"`

	// 公司名称
	CompanyName *string `json:"company_name,omitempty"`

	// 身份证明附件列表，至少1个，最多3个
	IdFiles []string `json:"id_files"`

	// 其他联系方式
	OtherWay *string `json:"other_way,omitempty"`
}

func (o ContactWayInfoV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ContactWayInfoV2 struct{}"
	}

	return strings.Join([]string{"ContactWayInfoV2", string(data)}, " ")
}
