### 介绍：
[介绍直达](https://github.com/zsr228/onenote-reminder/wiki/onenote-reminder介绍)

### 背景：   

### 目的：
	加载onenote的提醒事项进行提醒
### 方法：
	通过程序读取onenote的内容，然后进行事项提醒
### 功能：
	1、有提醒时，处理提醒
	2、有提醒时，延迟提醒(10min, 30min, 设置任意时间)
	3、有提醒时，删除提醒
	
	
### v1 feature list:   
    1、task达到datetime时，弹出框提醒;[done]   
    2、弹出框关闭时，触发close，再次达到1时，能弹出;[done 功能上已实现，还需优化]   
    3、增加desktop icon,默认后台运行，不占用任务栏;[以守护进程进行运行]
    4、添加定时读取笔记内容;[done 每隔5分钟读取笔记内容]
    5、将VanlliaAddIn设置为默认读取提醒事项的内容;[done]   
    6、在窗口增加强制刷新reminder内容;   
    7、设置主窗口隐藏;[done 需要点击关闭];   
    8、reminder弹出时,task背景颜色提醒;[fixme 目前是使用的bold]   
    9、已完成的task,不再显示;   


### v2 feature list:
    1、通过窗口来修改时间,同时更新到onenote中;   
    2、日志归档；   
    3、每周的一天能重复提醒;   

### bug list:
    1、重新读xml时,map数据没有清除,导致数据重复显示;[done]   
    2、日志log最好能中文显示，或者显示纯应为,不要magic log;   
    3、通知start的channel使用错误,导致channel阻塞,读取文件只执行两次;[done]   
    4、task修改时间后,task和datetime会重复显示;[done]   
    5、时分秒的task不显示;[done]   
    6、一分钟内显示两次提示;[done]   

### note list:
    1、v1.0使用表格形式的提醒事项;
