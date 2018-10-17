# TOOLS
# TOOLS/studygolang/Refresh_web/Refresh_web.go

<p>D:\Works\FishStoreLog\src>Refresh_web.exe </p>
<pre><code>Refresh_web.exe</code></pre>

iniPath: D:\Works\FishStoreLog\src\config.ini

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

请按任意键继续...

# goDate.exe
[goDate.go] (https://github.com/foxuc/TOOLS/blob/master/studygolang/goDate/goDate.go)

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


[回到顶部](#readme)
