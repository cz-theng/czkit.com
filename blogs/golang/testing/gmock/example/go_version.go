/**
* Author: CZ (cz.devnet@gmail.com)
*
* Description:
 */

package main

import (
	"github.com/cz-it/blog/blog/Go/testing/gomock/example/spider"
)

func GetGoVersion(s spider.Spider) string {
	body := s.GetBody()
	return body
}
