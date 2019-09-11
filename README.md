# qsuits-exec-go  
qsuits-java 代理执行工具 by golang  

|操作系统|程序名|地址|
|---|-----|---|
|windows 32 位|qsuits_windows_386.exe|[下载](https://github.com/NigelWu95/qsuits-exec-go/raw/master/bin/qsuits_windows_386.exe)|
|windows 64 位|qsuits_windows_amd64.exe|[下载](https://github.com/NigelWu95/qsuits-exec-go/raw/master/bin/qsuits_windows_amd64.exe)|
|linux 32 位|qsuits_linux_386|[下载](https://github.com/NigelWu95/qsuits-exec-go/raw/master/bin/qsuits_linux_386)|
|linux 64 位|qsuits_linux_amd64|[下载](https://github.com/NigelWu95/qsuits-exec-go/raw/master/bin/qsuits_linux_amd64)|
|mac 32 位|qsuits_darwin_386|[下载](https://github.com/NigelWu95/qsuits-exec-go/raw/master/bin/qsuits_darwin_386)|
|mac 64 位|qsuits_darwin_amd64|[下载](https://github.com/NigelWu95/qsuits-exec-go/raw/master/bin/qsuits_darwin_amd64)|

Usage:  
&ensp;&ensp;&ensp;&ensp;&ensp; this tool is a agent program for qsuits, your local environment need java8 or above. In default mode, this tool will use latest java qsuits to exec, you only need use qsuits-java's parameters to run. If you use local mode it mean you dont want to update latest qsuits automatically.  
Options:  
&ensp;&ensp;&ensp;&ensp;&ensp; -Local/-L &ensp;&ensp;&ensp;&ensp;&ensp;&ensp; use current default qsuits version to exec.  
&ensp;&ensp;&ensp;&ensp;&ensp; --help/-h/help &ensp;&ensp; print usage.  
Commands:  
&ensp;&ensp;&ensp;&ensp;&ensp; help &ensp;&ensp;&ensp;&ensp;&ensp;&ensp;&ensp;&ensp;&ensp;&ensp; print usage.  
&ensp;&ensp;&ensp;&ensp;&ensp; upgrade &ensp;&ensp;&ensp;&ensp;&ensp;&ensp;&ensp; upgrade this own executable program by itself.  
&ensp;&ensp;&ensp;&ensp;&ensp; versions &ensp;&ensp;&ensp;&ensp;&ensp;&ensp;&ensp; list all qsuits versions from local.  
&ensp;&ensp;&ensp;&ensp;&ensp; clear &ensp;&ensp;&ensp;&ensp;&ensp;&ensp;&ensp;&ensp;&ensp;&ensp; remove all old qsuits versions from local.  
&ensp;&ensp;&ensp;&ensp;&ensp; current &ensp;&ensp;&ensp;&ensp;&ensp;&ensp;&ensp;&ensp; query local default qsuits version.  
&ensp;&ensp;&ensp;&ensp;&ensp; chgver <no.> &ensp;&ensp;&ensp; set local default qsuits version.  
&ensp;&ensp;&ensp;&ensp;&ensp; download <no.> &ensp; download qsuits with specified version.  

Usage of qsuits-java:  https://github.com/NigelWu95/qiniu-suits-java  
