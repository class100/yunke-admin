package admin

import (
	`github.com/class100/yunke-core`
	`github.com/storezhang/gox`
)

// AddClient 添加客户端
func (oc *OrgClient) AddClient(client yunke.BaseClient, version yunke.ApiVersion) (c *yunke.BaseClient, err error) {
	c = new(yunke.BaseClient)
	orgClient := &struct {
		yunke.BaseClient

		// 原始未打包的文件编号
		OriginalFile string `json:"originalFile"`
	}{
		BaseClient:   client,
		OriginalFile: client.File,
	}
	// 清空原来的文件（防止提交的数据和机构数据定义有冲突）
	orgClient.File = ""

	err = oc.requestApi(yunke.OrgApiClientAddUrl, gox.HttpMethodPost, orgClient, nil, version, c)

	return
}
