package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ModifyStarRocksSecurityGroupReq struct {

	// 安全组ID。如果实例所选用的子网开启网络ACL进行访问控制，则该参数非必选。如果未开启ACL进行访问控制，则该参数必选。获取方法如下：  方法1：登录虚拟私有云服务的控制台界面，在安全组的详情页面查找安全组ID。  方法2：通过虚拟私有云服务的API接口查询，具体操作可参考查询安全组列表。
	SecurityGroupId string `json:"security_group_id"`
}

func (o ModifyStarRocksSecurityGroupReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyStarRocksSecurityGroupReq struct{}"
	}

	return strings.Join([]string{"ModifyStarRocksSecurityGroupReq", string(data)}, " ")
}
