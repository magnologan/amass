// Copyright 2017 Jeff Foley. All rights reserved.
// Use of this source code is governed by Apache 2 LICENSE that can be found in the LICENSE file.

package sources

import "github.com/OWASP/Amass/amass/core"

type OpenUKArchive struct {
	BaseDataSource
	baseURL string
}

func NewOpenUKArchive(srv core.AmassService) DataSource {
	o := &OpenUKArchive{baseURL: "http://www.webarchive.org.uk/wayback/archive"}

	o.BaseDataSource = *NewBaseDataSource(srv, core.ARCHIVE, "Open UK Arc")
	return o
}

func (o *OpenUKArchive) Query(domain, sub string) []string {
	if sub == "" {
		return []string{}
	}

	names, err := o.crawl(o.baseURL, domain, sub)
	if err != nil {
		o.Service.Config().Log.Printf("%v", err)
	}
	return names
}

func (o *OpenUKArchive) Subdomains() bool {
	return true
}
