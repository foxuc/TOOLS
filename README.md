# barcode_2wei
[barcode_2wei.exe](https://github.com/foxuc/TOOLS/blob/master/studygolang/barcode_2wei/barcode_2wei.exe)
```
C:\go>barcode_2wei.exe http://www.baidu.com

2018/10/19 10:47:08 ===  20181018 ver:0.1 by hi==

2018/10/19 10:47:08 Original data: http://www.baidu.com

2018/10/19 10:47:08 Encoded data:  http://www.baidu.com

2018/10/19 10:47:08 Out_barcode_2018_10_19_10_47_07.png
```
请按任意键继续. . .
输出百度www二维码：

![baidu.com](https://github.com/foxuc/TOOLS/blob/master/studygolang/barcode_2wei/Out_barcode_2018_10_19_10_47_07.png)



# TOOLS
# TOOLS/studygolang/Refresh_web/Refresh_web.go

<p>D:\Works\FishStoreLog\src>Refresh_web.exe </p>
<pre><code>Refresh_web.exe</code></pre>

iniPath: D:\Works\FishStoreLog\src\config.ini
```
          -------------- config.ini by xiaohai 2018.10.11 Ver:0.3 --------------
          
url:  http://www.baidu.com

cycles:  3

sleepMillisecond:  500

isPrint:  true

isPrintwebBody:  false

           -------------- config.ini --------------



2018-10-11 15:20:51.494423 [   Get.Url开始运行中...  ]: http://www.baidu.com

2018-10-11 15:20:51.549426 [Refresh_web OK]: http://www.baidu.com

2018-10-11 15:20:52.085457 [Refresh_web OK]: http://www.baidu.com

2018-10-11 15:20:52.620487 [Refresh_web OK]: http://www.baidu.com

ForNum: 9  sleepMillisecond: 500ms  Use: 4.3262474s
 ```   
请按任意键继续...
       
# 【方法1】goDate.exe 设置BAT几天前
[goDate.go] (https://github.com/foxuc/TOOLS/blob/master/studygolang/goDate/goDate.go)
```
D:\Works\>goDate.exe +10

20181027

D:\Works\>goDate.exe -1

20181016

D:\Works\>goDate.exe -14

20181003

D:\Works\>goDate.exe +

Parameter Error !

D:\Works\>goDate.exe -

Parameter Error !

D:\Works\>goDate.exe -0

Parameter Error !

D:\Works\>goDate.exe +0

Parameter Error !
```

# Run goDate.exe To BAT
```
echo off
set runexe=goDate.exe -14
%runexe% > 1.txt

::skip=1 读取第二行的内容

for /f "delims=" %%a in (1.txt) do (
set txt=%%~a
)

echo on
echo %runexe% "day:" %txt%

if exist ".\BAK%txt%" (
    rmdir /S /Q .\BAK%txt%
)
mkdir .\BAK%txt%
pause
```
>>>
D:\Works\>echo off

D:\Works\>echo goDate.exe -14 "day:" 20181003

goDate.exe -14 "day:" 20181003

D:\Works\>if exist ".\BAK20181003" (rmdir /S /Q .\BAK20181003 )

D:\Works\>mkdir .\BAK20181003

D:\Works\>pause

请按任意键继续. . .
>>>


# 【方法2】Vbs 设置BAT几天前
```
pushd F:\BAKVM\
set n=14

>"%tmp%\t.vbs" echo;t=date()-%n%:y=right(year(t),4):m=right("0"^&month(t),2):d=right("0"^&day(t),2):wscript.echo y^&" "^&m^&" "^&d

for /f "tokens=1-3" %%a in ('cscript /nologo "%tmp%\t.vbs"') do set y=%%a&set m=%%b&set d=%%c

echo;%n%Date:%y%%m%%d%
echo;"F:\BAKVM\BAK%y%%m%%d%"
if exist "F:\BAKVM\BAK%y%%m%%d%" (
    rmdir /S /Q F:\BAKVM\BAK%y%%m%%d%
)
::几天前--
```

# 【方法3】powershell 删除7天前的日志
```
把以下命令保存为ps1脚本或加到计划任务中：
#delete logs in specify website, just save logs in eight days~   
$TimeOutDays=7
$filePath="C:\temp\"    
$allFiles=get-childitem -path $filePath
 
foreach ($files in $allFiles)    
{      
   $daypan=((get-date)-$files.lastwritetime).days      
   if ($daypan -gt $TimeOutDays)      
   { 
     #$files.FullName
     remove-item $files.fullname -Recurse -force      
    }    
}
```
参数说明:
-Recurse  表示递归，删除子文件和子文件夹

-Force 表示强制删除，不询问

[回到顶部](#readme)
