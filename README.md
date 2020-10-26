# mosquito
## 背景
 对比目前市面所有文档系统，发现功能过于单一，对公司的共享目录文档管理方式无法兼容，共享文档安全性低
 同时互联网对文档管理更加偏向文本化（markdown,html,docx），更加容易在各平台系统兼容。
 
目前市面上文档管理系统比较：
 - wookteam：支持在线文档（markdown+流程图），不支持其他文档，不支持备份导出转换，[文档](https://gitee.com/aipaw/wookteam/blob/master/install/DOCKER.md)。
 - wiki(confluence+MediaWiki+Docsify+vuepress)：支持富客户端+markdown，不支持其他文档管理[文档](https://www.jianshu.com/p/f79236289793)。
 - showdoc:支持在线接口api定制，markdown格式，无法支持其他文档[文档](https://www.showdoc.com.cn/demo?page_id=7)。
 
## 简介
mosquito是一款以文件系统作为基础的在线文档管理系统，
在线文档管理系统功能包括：
- 基于目录树的在线文档查看
- office办公套件预览（编辑暂不支持）
- pdf在线预览
- html/文本编辑器
- 在线代码编辑器
- 图片预览功能（后续支持在线绘图）
- markdown编辑器
- 思维导图编辑器
- 在线作图编辑器（流程图 ，活动图，类图，时序图等）

 ## 系统架构
 ![系统拓扑图](doc/topology.png)
 ### 前端依赖插件
 - pdfjs（pdf预览器）[教程](http://mozilla.github.io/pdf.js/)
 - mavon-editor（markdown编辑器）[教程](https://github.com/hinesboy/mavonEditor)
 - codemirror（代码编辑工具）[教程](https://github.com/surmon-china/vue-codemirror)
 - wangeditor（html文本编辑器）[教程](http://www.wangeditor.com/)
 - x-data-spreadsheet（excel预览工具）[教程](https://github.com/myliang/x-spreadsheet)
 - mxgraph(废弃，流程图，无导出功能) [教程](https://jgraph.github.io/mxgraph/javascript/examples/grapheditor/www/index.html)
 - 乐吾乐 Topology（正方形错误，其他图形待绘制）[教程](https://www.yuque.com/alsmile/topology/make-shape)
 - kityminder-core（思维导图）[教程](https://github.com/fex-team/kityminder-core/wiki/command)
### 后端依赖插件
 - beego [教程](https://beego.me/docs/intro/)
 - go-smb2 [教程](https://github.com/hirochachacha/go-smb2)
 - goftp [教程](https://github.com/dutchcoders/goftp)
 ## 版本更新
 1. 1.0.0：版本支持全局的文档管理基建。
 2. 1.0.1：权限+个人文档在+在线photoshop作图（暂未开发）。
 ## 相关文档
 1. [用户操作手册](doc/user.md)
 2. [运维手册](doc/oper.md)
 3. [二次开发引导手册](doc/dev.md) <br/>
 3.1. [前端关联编辑器](doc/devf.md)<br/>
 3.2. [后端文件系统开发](doc/devb.md)
 4. [参数配置手册](doc/conf.md)
 