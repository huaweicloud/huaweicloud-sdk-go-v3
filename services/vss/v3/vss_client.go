package v3

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/vss/v3/model"
)

type VssClient struct {
	HcClient *http_client.HcHttpClient
}

func NewVssClient(hcClient *http_client.HcHttpClient) *VssClient {
	return &VssClient{HcClient: hcClient}
}

func VssClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

//认证租户的域名资产
func (c *VssClient) AuthorizeDomains(request *model.AuthorizeDomainsRequest) (*model.AuthorizeDomainsResponse, error) {
	requestDef := GenReqDefForAuthorizeDomains()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AuthorizeDomainsResponse), nil
	}
}

//创建租户的域名资产
func (c *VssClient) CreateDomains(request *model.CreateDomainsRequest) (*model.CreateDomainsResponse, error) {
	requestDef := GenReqDefForCreateDomains()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDomainsResponse), nil
	}
}

//删除租户的域名资产
func (c *VssClient) DeleteDomains(request *model.DeleteDomainsRequest) (*model.DeleteDomainsResponse, error) {
	requestDef := GenReqDefForDeleteDomains()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteDomainsResponse), nil
	}
}

//获取租户的所有域名资产
func (c *VssClient) ListDomains(request *model.ListDomainsRequest) (*model.ListDomainsResponse, error) {
	requestDef := GenReqDefForListDomains()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDomainsResponse), nil
	}
}

//获取域名登录配置
func (c *VssClient) ShowDomainSettings(request *model.ShowDomainSettingsRequest) (*model.ShowDomainSettingsResponse, error) {
	requestDef := GenReqDefForShowDomainSettings()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDomainSettingsResponse), nil
	}
}

//更新域名登录配置
func (c *VssClient) UpdateDomainSettings(request *model.UpdateDomainSettingsRequest) (*model.UpdateDomainSettingsResponse, error) {
	requestDef := GenReqDefForUpdateDomainSettings()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateDomainSettingsResponse), nil
	}
}

//获取域名业务风险扫描结果
func (c *VssClient) ListBusinessRisks(request *model.ListBusinessRisksRequest) (*model.ListBusinessRisksResponse, error) {
	requestDef := GenReqDefForListBusinessRisks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListBusinessRisksResponse), nil
	}
}

//获取域名端口扫描结果
func (c *VssClient) ListPortResults(request *model.ListPortResultsRequest) (*model.ListPortResultsResponse, error) {
	requestDef := GenReqDefForListPortResults()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPortResultsResponse), nil
	}
}

//获取域名漏洞扫描结果
func (c *VssClient) ShowResults(request *model.ShowResultsRequest) (*model.ShowResultsResponse, error) {
	requestDef := GenReqDefForShowResults()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowResultsResponse), nil
	}
}

//更新域名扫描漏洞的误报状态
func (c *VssClient) UpdateFalsePositive(request *model.UpdateFalsePositiveRequest) (*model.UpdateFalsePositiveResponse, error) {
	requestDef := GenReqDefForUpdateFalsePositive()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateFalsePositiveResponse), nil
	}
}

//取消或重启域名漏洞扫描任务
func (c *VssClient) CancelTasks(request *model.CancelTasksRequest) (*model.CancelTasksResponse, error) {
	requestDef := GenReqDefForCancelTasks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CancelTasksResponse), nil
	}
}

//创建域名漏洞扫描任务并启动
func (c *VssClient) CreateTasks(request *model.CreateTasksRequest) (*model.CreateTasksResponse, error) {
	requestDef := GenReqDefForCreateTasks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateTasksResponse), nil
	}
}

//获取域名漏洞扫描的历史扫描任务
func (c *VssClient) ListTaskHistories(request *model.ListTaskHistoriesRequest) (*model.ListTaskHistoriesResponse, error) {
	requestDef := GenReqDefForListTaskHistories()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTaskHistoriesResponse), nil
	}
}

//获取域名漏洞扫描任务详情
func (c *VssClient) ShowTasks(request *model.ShowTasksRequest) (*model.ShowTasksResponse, error) {
	requestDef := GenReqDefForShowTasks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTasksResponse), nil
	}
}
