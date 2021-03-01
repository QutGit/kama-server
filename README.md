#  golang-started

教师线  golang started

### 使用echo

现在 echo 分支代码为 echo 框架实现，建议新项目使用echo

### ⚠️需要认真阅读并按此操作
- git clone git@gitlab.yc345.tv:teacher/backend/golang-started.git 新项目名称(如 teacher-courseware)
- git remote add template git@gitlab.yc345.tv:teacher/backend/golang-started.git
- git remote rm origin
- git remote add origin 新项目地址

### 配置文件

config 文件夹中的配置文件根据GO_ENV 进行加载，默认加载local(本地开发)
local 文件不再版本控制中，只用于本地

### gitlab-ci

使用gitlab-ci 直接部署到服务，ci 中 PORT 需要根据实际情况进行修改

### courseware

- controller 控制器
- model 模型
- router 路由
- service 操作函数

### TODO
- 单元测试
- swagger 自动生成

