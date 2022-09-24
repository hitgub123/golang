        bug：
            在项目根目录go mod init 任意名称 后 go mod tidy后， 
            a1.go里import "github.com/gin-gonic/gin"处红线并报错：
            could not import github.com/gin-gonic/gin (cannot find package "github.com/gin-gonic/gin" in any of**** ，
            且用go run .\a1.go运行gin的demo代码后马上停止(没报错)，浏览器无法访问项目。
            一开始按提示去GOROOT和GOPATH下找，发现gin-gonic下没有gin，只有gin@v1.8.1，以为是这个原因。其实跟这个无关。
            其实只要正确init和tidy就没问题。项目启动失败原因是demo代码有问题，修改代码后即可启动项目。
            但文件内部依旧有红线并报错，vscode提示Error loading workspace: You are outside of a module**** ，
            原因是自己vscode打开的folder是项目的上一级目录，改为打开项目根目录即可解决。
