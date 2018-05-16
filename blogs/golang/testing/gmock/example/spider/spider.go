/**
* Author: CZ (cz.devnet@gmail.com)
*
* Description:
 */

package spider

//go:generate mockgen -destination mock_spider.go -package spider github.com/cz-it/czkit.com/blogs/golang/testing/gmock/example/spider Spider

type Spider interface {
	GetBody() string
}
