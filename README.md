这是go module测试项目，测试使用不同版本的包

1 :  语义化版本
    语义化版本号规范把版本号如 v1.2.3 中的 1 定义为主要版本号，2 为次要版本号，3 为修订版本号 )

2 ： 在仓库创建tag标签

    git tag v1.0.0
    git push --tags

3 ： go mod 命令

    go mod init:初始化modules
    go mod download:下载modules到本地cache
    go mod edit:编辑go.mod文件，选项有-json、-require和-exclude，可以使用帮助go help mod edit
    go mod graph:以文本模式打印模块需求图
    go mod tidy:删除错误或者不使用的modules
    go mod vendor:生成vendor目录
    go mod verify:验证依赖是否正确
    go mod why：查找依赖

4 ： go get更新module
    运行 go get -u 将会升级到最新的次要版本或者修订版本
    运行 go get -u=patch 将会升级到最新的修订版本（比如说，将会升级到 1.0.1 版本，但不会升级到 1.1.0 版本）
    运行 go get package@version将会升级到指定的版本号
