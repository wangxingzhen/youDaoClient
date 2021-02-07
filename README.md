## 有道词典命令行版客户端

(服务端代码会在后期开放出来，目前还在优化中)

### 逻辑
当前程序分为服务端+客户端，服务端主要负责与有道文本翻译接口进行交互，客户端主要与服务端交互

### 简单说明
打包命令
`go build`
生成可执行文件，然后把可执行文件放到系统全局变量里面即可

例子：

```bigquery
[root@localhost youDaoClient]# fanyi 翻译一下
待翻译的文本：
        翻译一下
翻译结果：
        translation
网络翻译：
        翻译一下：south of the slot; traduce;
        人工翻译一下：Snow Halation;
        翻译一下这句英语：hope you can understand me;
```
### 下载地址
（发现一个已知问题：在win10原生命令行下，因为编码格式的原因，输出的汉语会乱码，推荐使用cmder地址https://cmder.net/）

win客户端
[下载](https://gitee.com/wang-xingzhen/you-dao-client/raw/master/exec/win/fanyi.exe "下载")

linux客户端
[下载](https://gitee.com/wang-xingzhen/you-dao-client/raw/master/exec/linux/fanyi.zip "下载")

osx客户端
[下载](https://gitee.com/wang-xingzhen/you-dao-client/raw/master/exec/osx/fanyi.zip "下载")

### 使用说明
下载对应的客户端之后，可以配置一下系统全局变量，这样在任何一个路径下都可以使用了，需要注意的是，需要联网，且服务器在北京的阿里云上，请知晓