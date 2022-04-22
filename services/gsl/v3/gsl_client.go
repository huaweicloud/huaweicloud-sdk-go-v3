package v3

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/gsl/v3/model"
)

type GslClient struct {
	HcClient *http_client.HcHttpClient
}

func NewGslClient(hcClient *http_client.HcHttpClient) *GslClient {
	return &GslClient{HcClient: hcClient}
}

func GslClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

// 查询套餐列表信息
//
// 查询套餐列表信息
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *GslClient) ListProPricePlans(request *model.ListProPricePlansRequest) (*model.ListProPricePlansResponse, error) {
	requestDef := GenReqDefForListProPricePlans()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListProPricePlansResponse), nil
	}
}

// 清除实名认证信息
//
// 清除实名认证信息，接口只支持电信卡调用
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *GslClient) DeleteRealName(request *model.DeleteRealNameRequest) (*model.DeleteRealNameResponse, error) {
	requestDef := GenReqDefForDeleteRealName()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteRealNameResponse), nil
	}
}

// 激活实体卡
//
// 创建激活实体卡申请，返回业务受理单号。1~2个工作日完成激活操作。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *GslClient) EnableSimCard(request *model.EnableSimCardRequest) (*model.EnableSimCardResponse, error) {
	requestDef := GenReqDefForEnableSimCard()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.EnableSimCardResponse), nil
	}
}

// 查询SIM卡列表
//
// 查询SIM卡列表
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *GslClient) ListSimCards(request *model.ListSimCardsRequest) (*model.ListSimCardsResponse, error) {
	requestDef := GenReqDefForListSimCards()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSimCardsResponse), nil
	}
}

// SIM卡机卡重绑
//
// 支持固定机卡重绑(需要上传IMEI，将SIM卡绑定到指定IMEI的设备)和普通机卡重绑(会清除之前绑定的设备,将SIM卡绑定到正在使用的设备)，单卡每月只允许重绑2次，接口只支持电信卡调用。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *GslClient) RegisterImei(request *model.RegisterImeiRequest) (*model.RegisterImeiResponse, error) {
	requestDef := GenReqDefForRegisterImei()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RegisterImeiResponse), nil
	}
}

// SIM卡单卡复机
//
// 创建复机申请，返回业务受理单号。1~2个工作日完成复机操作。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *GslClient) ResetSimCard(request *model.ResetSimCardRequest) (*model.ResetSimCardResponse, error) {
	requestDef := GenReqDefForResetSimCard()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ResetSimCardResponse), nil
	}
}

// SIM卡达量断网/恢复在用
//
// SIM卡达量断网/恢复在用,只支持电信卡。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *GslClient) SetExceedCutNet(request *model.SetExceedCutNetRequest) (*model.SetExceedCutNetResponse, error) {
	requestDef := GenReqDefForSetExceedCutNet()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SetExceedCutNetResponse), nil
	}
}

// 实体卡限速
//
// 实体卡限速接口,支持电信和联通实体卡调用。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *GslClient) SetSpeedValue(request *model.SetSpeedValueRequest) (*model.SetSpeedValueResponse, error) {
	requestDef := GenReqDefForSetSpeedValue()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SetSpeedValueResponse), nil
	}
}

// 查询SIM卡实名认证信息
//
// 实时查询SIM卡实名认证信息。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *GslClient) ShowRealNamed(request *model.ShowRealNamedRequest) (*model.ShowRealNamedResponse, error) {
	requestDef := GenReqDefForShowRealNamed()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRealNamedResponse), nil
	}
}

// 查询SIM卡详情
//
// 查询SIM卡详情
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *GslClient) ShowSimCard(request *model.ShowSimCardRequest) (*model.ShowSimCardResponse, error) {
	requestDef := GenReqDefForShowSimCard()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSimCardResponse), nil
	}
}

// SIM卡申请断网/恢复在用
//
// SIM卡申请断网/恢复在用,只支持电信卡。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *GslClient) StartStopNet(request *model.StartStopNetRequest) (*model.StartStopNetResponse, error) {
	requestDef := GenReqDefForStartStopNet()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StartStopNetResponse), nil
	}
}

// SIM卡单卡停机
//
// 创建停机申请，返回业务受理单号。1~2个工作日完成停机操作。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *GslClient) StopSimCard(request *model.StopSimCardRequest) (*model.StopSimCardResponse, error) {
	requestDef := GenReqDefForStopSimCard()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StopSimCardResponse), nil
	}
}

// 查询流量池成员列表
//
// 查询流量池成员列表
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *GslClient) ListSimPoolMembers(request *model.ListSimPoolMembersRequest) (*model.ListSimPoolMembersResponse, error) {
	requestDef := GenReqDefForListSimPoolMembers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSimPoolMembersResponse), nil
	}
}

// 查询流量池列表
//
// 查询流量池列表
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *GslClient) ListSimPools(request *model.ListSimPoolsRequest) (*model.ListSimPoolsResponse, error) {
	requestDef := GenReqDefForListSimPools()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSimPoolsResponse), nil
	}
}

// 批量查询实体卡流量
//
// 批量查询实体卡流量
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *GslClient) ListFlowBySimCards(request *model.ListFlowBySimCardsRequest) (*model.ListFlowBySimCardsResponse, error) {
	requestDef := GenReqDefForListFlowBySimCards()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListFlowBySimCardsResponse), nil
	}
}

// sim卡套餐列表查询
//
// SIM卡套餐列表查询，实体卡只会有一个套餐，eSIM/vSIM可能会有多个套餐
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *GslClient) ListSimPricePlans(request *model.ListSimPricePlansRequest) (*model.ListSimPricePlansResponse, error) {
	requestDef := GenReqDefForListSimPricePlans()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSimPricePlansResponse), nil
	}
}
