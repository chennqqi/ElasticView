 <h1 align="center">
   <br>
   <img src="https://gitee.com/cynthia520/elastic-view/raw/main/static/show_img/2.png"/>
   <br>
   ElasticView
   <br>
</h1>

-----------
[![Go Documentation](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](https://godoc.org/github.com/1340691923/ElasticView)
[![license](https://img.shields.io/github/license/mashape/apistatus.svg?maxAge=2592000)](https://github.com/1340691923/ElasticView/blob/main/LICENSE)
[![Release](https://img.shields.io/github/release/1340691923/ElasticView.svg?label=Release)](https://gitee.com/cynthia520/elastic-view/releases)

> ElasticView
> 是一款用来监控ElasticSearch状态和操作ElasticSearch索引的web可视化工具。它由golang开发而成，具有部署方便，占用内存小等优点，官网地址:http://www.elastic-view.cn

* ElasticSearch连接树管理（更方便的切换测试/生产环境）
* 支持权限管理
* 支持sql转换成dsl语法
* 更方便的重建索引
* 任务管理
* 备份管理
* 可将查询内容下载为excel文件
* 可进行索引创建，映射创建，别名创建，索引删除等操作
* 支持版本 `6.x`,`7.x`,`8.x`
* 支持类似Navicat功能
* docker部署
* 支持sqlite3（免安装gcc版）
* 数据抽取功能

## 官网地址

[http://www.elastic-view.cn]( http://www.elastic-view.cn)

## Quick Start

1. [应用程序下载地址]( https://gitee.com/cynthia520/elastic-view/releases/)
2. 下载应用程序下载地址里面的对应压缩包后解压（windows用户下载ElasticView_windows.zip，linux用户下载ElasticView_linux.zip，mac用户下载
   ElasticView_mac.zip）
3. （若无需mysql存储数据则跳过该步骤）修改config.json文件中的 数据库连接信息，日志存放目录和应用启动端口等配置信息
4. （若无需mysql存储数据则跳过该步骤）数据存储 若config.json下的dbType为sqlite3则指定sqlite配置下的dbPath即可（无需安装gcc）
   为mysql则需新建mysql数据库 es_view，导入es_view.sql并修改mysql配置
5. windows：双击ElasticView.exe linux：chmod +x ElasticView && nohup ./ElasticView > ElasticView.log &
6. 浏览器访问对应ip:端口，初始用户名：admin，初始密码：admin

## ElasticView 教学视频（来源：https://blog.csdn.net/dzcro/article/details/126918571?spm=1001.2014.3001.5502）

https://www.bilibili.com/video/av473231184/?zw&spm_id_from=888.80996.embed_old

## ElasticView 教学视

##支持操作系统：

- Windows
- Linux
- MacOs

<!-- ## 手动编译
 1. `git clone git@github.com:1340691923/ElasticView.git`
 2. `cd vue && cnpm install (安装前端依赖)`
 3. `执行static/build 下的vue_build.bat （打前端正式包）`
 4. `执行static/build 下的win_build.bat(linux则为linux_build.bat)  (编译二进制可执行程序)`
  -->

## ☁docker部署

图文教程：[https://4xx.me/archives/128.html](https://4xx.me/archives/128.html)

- sqlite3存储

```shell
docker run --restart=unless-stopped -d -p 8090:8090 -v /data/elastic_view/data:/data -v /data/elastic_view/logs:/logs kecikeci/elastic_view:latest
```

- mysql存储部署

1. 新建数据库`es_view`
2. 导入对应数据表`es_view.sql`
3. 准备`config.yml`配置文件并修改对应连接信息

```yaml
log:
   storageDays: 4          # 日志保留天数
   logDir: "logs"         # 日志保留文件夹
port: 8090              # 启动端口
dbType: "mysql"       # 数据保留类型 分为 sqlite3 和 mysql
sqlite:                 # dbType为sqlite3时填 dbPath为数据保存文件地址
   dbPath: "es_view.db"
mysql:                  # dbType为mysql时填
   username: "root"
   pwd: "123456"
   ip: "127.0.01"
   port: "3306"
   dbName: "es_view"
   maxOpenConns: 10
   maxIdleConns: 10
esPwdSecret: "concat_mail!!->1340691923@qq.com" # es密码加密密钥 加密方式为 AES
version: "1.8.5"  # ES 版本号
deBug: false      # 是否为测试模式 如果为 false则打开默认浏览器直接访问地址
```

4. 挂载`config`目录或`config.yml`配置文件启动镜像

```shell
docker run --restart=unless-stopped -d -p 8090:8090 -v /data/elastic_view/data:/data -v /data/elastic_view/config:/config -v /data/elastic_view/logs:/logs kecikeci/elastic_view:latest
```

- 浏览器访问对应ip:8090，初始用户名：admin，初始密码：admin

## 🛠️手动构建

```shell
# 拉取项目源代码
git clone https://github.com/1340691923/ElasticView

# 同步前端项目依赖
cd resources/vue && npm install

# 构建前端包
npm run build

# 构建项目二进制程序
go build -o ev cmd/ev
```

## 更多

* 铸龙用户行为分析系统     https://github.com/1340691923/xwl_bi
* 软考成绩快查工具        https://github.com/1340691923/SoftTestMonitor

### 求职中，个人微信二维码

<!-- ![image](https://gitee.com/cynthia520/elastic-view/raw/main/static/show_img/weixin.jpg)   -->
<img src="https://raw.githubusercontent.com/1340691923/ElasticView/main/resources/show_img/weixin.jpg" style="width: 220px"> 
### 求职中，团队成员微信二维码

<img src="http://cdn.ycvod.net/img/wx2.jpg" style="width: 220px"> 

### QQ群

<img src="https://raw.githubusercontent.com/1340691923/ElasticView/main/resources/show_img/qq_group.jpg" style="width: 220px">
