package yunke

import (
	`strconv`

	`github.com/class100/yunke-core`
	`github.com/storezhang/gox`
)

// GetOrg 取得机构的信息
func (a *Admin) GetOrg() (org *Org, err error) {
	org = new(Org)
	err = a.request(yunke.AdminApiGetOrgUrl, gox.HttpMethodGet, nil, map[string]string{
		"id": strconv.FormatInt(a.Id, 10),
	}, org)

	return
}
