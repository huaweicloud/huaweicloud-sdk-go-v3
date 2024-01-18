package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CbcFreezeInfo struct {

	// 在冻结/解冻等操作下，云服务需要达到的主要效果：  - 1：（实现/去除）冻结可释放。（资源冻结后，客户可以手动删除/释放云资源和云资源上数据。）  - 2：（实现/去除）冻结不可释放。（资源冻结后，客户不能手动删除/释放云资源以及云资源上数据，相当于云服务被贴了封条，不能改变数据和资源。对应解冻后，就可以删除和修改客户数据了。）  - 3：（实现/去除）冻结后不可续费。（资源冻结后，资源不能发起续费操作；解冻后，才可以发起续费操作。）  - effect字段和上面status字段(1冻结、0解冻)配合使用，表示在发起冻结/解冻命令下，云服务达到的冻结效果。  - 为空时，默认为effect=1（云服务需要能兼容处理，默认当做effect=1）。  - 注：云服务是根据status和effect在真实限制云服务的操作/API等。不是使用下文的scene字段去做云服务操作/API的限制。下文的scene字段，主要用于Console页面的tips、API错误码等客户体验使用。
	Effect *int32 `json:"effect,omitempty"`

	Scene *CbcFreezeScene `json:"scene,omitempty"`
}

func (o CbcFreezeInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CbcFreezeInfo struct{}"
	}

	return strings.Join([]string{"CbcFreezeInfo", string(data)}, " ")
}
