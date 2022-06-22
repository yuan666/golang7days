要使用CGO特性
    需要安装C/C++构建工具链，在macOS和Linux下是要安装GCC,在windows下是需要安装MinGW工具。
    保证环境变量 CGO_ENABLED 被设置为1，这表示CGO是被启用的状态。

2.2.1  import "C"
import "C"导入语句需要单独一行，不能与其他包一同import。

2.2.2 #cgo 语句
    // #cgo CFLAGS: -DPNG_DEBUG=1 -I./include
    // #cgo LDFLAGS: -L/usr/local/lib -lpng
    // #include <png.h>
    import "C"

2.2.3 build tag 条件编译
