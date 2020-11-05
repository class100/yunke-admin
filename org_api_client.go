package admin

import (
	`github.com/class100/yunke-core`
	`github.com/storezhang/gox`
)

// AddClient 添加客户端
func (co *ClientOrg) AddClient(client core.BaseClient, version core.ApiVersion) (c *core.BaseClient, err error) {
	c = new(core.BaseClient)
	orgClient := &struct {
		core.BaseClient

		// 原始未打包的文件编号
		OriginalFile string `json:"originalFile"`
	}{
		BaseClient:   client,
		OriginalFile: client.File,
	}
	// 清空原来的文件（防止提交的数据和机构数据定义有冲突）
	orgClient.File = ""

	err = co.requestApi(core.OrgApiClientAddUrl, gox.HttpMethodPost, orgClient, nil, version, c)

	return
}
