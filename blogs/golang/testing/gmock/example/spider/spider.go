/**
* Author: CZ (cz.devnet@gmail.com)
*
* Description:
 */

package spider

//go:generate mockgen -destination mock_spider.go -package spider github.com/cz-it/blog/blog/Go/testing/gomock/example/spider Spider

type Spider interface {
	GetBody() string
}
