package admin

import (
	`strconv`

	`github.com/class100/yunke-core`
	`github.com/storezhang/gox`
)

// GetOrg 取得机构的信息
func (ca *ClientAdmin) GetOrg() (org *core.Org, err error) {
	org = new(core.Org)
	err = ca.request(core.AdminApiGetOrgUrl, gox.HttpMethodGet, nil, map[string]string{
		"id": strconv.FormatInt(ca.Id, 10),
	}, org)

	return
}
