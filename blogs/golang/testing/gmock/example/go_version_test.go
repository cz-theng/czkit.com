/**
* Author: CZ (cz.devnet@gmail.com)
*
* Description:
 */

package main

import (
	"github.com/cz-it/blog/blog/Go/testing/gomock/example/spider"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestGetGoVersion(t *testing.T) {
	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()
	mockSpider := spider.NewMockSpider(mockCtl)
	mockSpider.EXPECT().GetBody().Return("go1.8.3")
	goVer := GetGoVersion(mockSpider)

	if goVer != "go1.8.3" {
		t.Error("Get wrong version %s", goVer)
	}
}
