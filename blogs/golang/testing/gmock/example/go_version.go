/**
* Author: CZ (cz.devnet@gmail.com)
*
* Description:
 */

package main

import (
	"github.com/cz-it/czkit.com/blogs/golang/testing/gmock/example/spider"
)

func GetGoVersion(s spider.Spider) string {
	body := s.GetBody()
	return body
}
