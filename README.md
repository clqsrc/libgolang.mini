# libgolang.mini
golang https/libs to so/dll for qt5/delphi/c/lazarus

# [cn]中文说明

本来这个项目名应该直接叫 libgolang 的，我好几年前就一直在工作中用着了。不过既然正式放到 github 上开源还是选择一个正式点的名称为好，毕竟叫 libgolang 的话，那事情也未必太大了。

另外，我也没打算所有的 go 语言都写成 dll 给其他地方用。实际上不但没有，我还从工作中的代码裁减掉了一部分，因为在工作中讲究实用，有些代码写得太过复杂，我自己后头看都不太想维护。所以打算分成两个目录，一个是 "full" 表示完整版本，一个是 "mini" 基本上就是放些 https tls js 等比较通用的功能。

基本上只放了 windows 环境的文件，其实我是在 ubuntu 22.04 也为主的。实际上两者也通用，其中涉及到的 cgo 部分编译器配置本来是用的另外的独立安装版本，不过因为工作中环境都装有 qt5 ，所以直接换上了 qt5 的 32 和 64 位编译器，倒也省了不少麻烦。不想装 qt5 的网友自己换 c/c++ 编译器，然后改在编译脚本的路径就行了。

再另外，我其实也是在生成 dll/so 文件时才运行一个那个简单到不行的脚本，其实大多娄代码测试运行时是用的 liteide 的配置切换，因为非常的方便。而写代码原来主要也用 liteide ，不过功能毕竟少，所以后期也用了不少 vscode 。把 vscode 配置好以后确实编写代码方便许多，但 vscode 还不能全部替代 liteide ，推荐大伙两个都装。


mini/delphi7 中是调用这个 so/dll 的 d7 版本示例。其实我主要用在 qt5 中，不过因为工作很忙，而我用 d7 感觉最方便，所以顺手就用它写了个例子，其实是想用我魔改过的 lazarus 写例子的。不过真心 delphi7 是世界上最好用的 ide -- 虽然因为形势比人强我已经很久没有用它了。

这个库虽然带有 "mini" 的名称，其实有些非常重要的小技巧，否则想用 go 来写库还要方便用，还是很困难的。容我以后慢慢道来。代码按我的想法精简过，工作其实是不够用的，不过也是希望“授人以渔”大家按自己的需求来扩充这个库自己 fork 份自己玩是最好。

