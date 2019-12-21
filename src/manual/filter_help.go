package manual

import "fmt"

func FilterUsage()  {

	fmt.Println("从数据源输入的数据通常可能存在过滤需求，如过滤指定规则的文件名、过滤时间点或者过滤存储类型等，可通过配置选项设置一些过滤条件，目前支持两种过滤条件：基本字段过滤和特殊特征匹配过滤。")
	fmt.Println()
	fmt.Println("1. 基本字段过滤")
	fmt.Println("根据设置的字段条件进行筛选，多个条件时需同时满足才保留，若存在记录不包该字段信息时则正向规则下不保留，反正规则下保留，字段包含：")
	fmt.Println("f-prefix=        表示选择文件名符合该前缀的文件")
	fmt.Println("f-suffix=        表示选择文件名符合该后缀的文件")
	fmt.Println("f-inner=         表示选择文件名包含该部分字符的文件")
	fmt.Println("f-regex=         表示选择文件名符合该正则表达式的文件，所填内容必须为正则表达式")
	fmt.Println("f-mime=          表示选择符合该 mime 类型（也即 content-type）的文件")
	fmt.Println("f-type=          表示选择符合该存储类型的文件，参下述关于 f-type|")
	fmt.Println("f-status=        表示选择符合该存储状态的文件, 为 0（启用） 或 1（禁用）")
	fmt.Println("f-date-scale     设置过滤的时间范围，格式为 [<date1>,<date2>]，<date> 格式为：yyyy-MM-DD HH:MM:SS，可参考：特殊规则")
	fmt.Println("f-anti-prefix=   表示排除文件名符合该前缀的文件")
	fmt.Println("f-anti-suffix=   表示排除文件名符合该后缀的文件")
	fmt.Println("f-anti-inner=    表示排除文件名包含该部分字符的文件")
	fmt.Println("f-anti-regex=    表示排除文件名符合该正则表达式的文件，所填内容必须为正则表达式")
	fmt.Println("f-anti-mime=     表示排除该 mime 类型的文件")
	fmt.Println()
	fmt.Println("（1）关于 f-type")
	fmt.Println("存储源  type 参数类型              具体值")
	fmt.Println("七牛    整型            0 表示标准存储；1 表示低频存储")
	fmt.Println("其他   字符串             如：Standard 表示标准存储")
	fmt.Println("（2）特殊字符")
	fmt.Println("特殊字符包括: , \\ = 如有参数值本身包含特殊字符需要进行转义：\\, \\\\ \\=")
	fmt.Println(" ")
	fmt.Println("（3）f-date-scale")
	fmt.Println("<date> 中的 00:00:00 为默认值可省略，无起始时间则可填 [0,<date2>]，结束时间支持 now 和 max，分别表示到当前时间为结束或无结束时间。如果使用命令行来设置，注意日期值包含空格的情况（date 日期和时刻中间含有空格分隔符），故在设置时需要使用引号 ' 或者 \"，如 -f-date-scale=\"[0,2018-08-01 12:30:00]\"，配置文件则不需要引号。")
	fmt.Println(" ")
	fmt.Println("2. 特殊特征匹配过滤 f-check")
	fmt.Println("根据资源的字段关系选择某个特征下的文件，目前支持 ext-mime 检查，程序内置的默认特征配置见：check 默认配置 https://github.com/NigelWu95/qiniu-suits-java/blob/master/resources/check.json，运行 参数选项如下：")
	fmt.Println("f-check=ext-mime 表示进行后缀名 ext 和 mimeType（即 content-type）匹配性检查，不符合规范的疑似异常文件将被筛选出来")
	fmt.Println("f-check-config 自定义资源字段规范对应关系列表的配置文件，格式为 json，自定义规范配置 key 字段必填，其元素类型为列表 [], 否则无效，如 ext-mime 配置时后缀名和 mimeType 用 : 组合成字符串成为一组对应关系，写法如下：")
	fmt.Println("{")
	fmt.Println("  \"ext-mime\": [")
	fmt.Println("    \"mp5:video/mp5\"")
	fmt.Println("  ]")
	fmt.Println("}")

	fmt.Println("f-check-rewrite 是否覆盖默认的特征配置，为 false（默认）表示将自定义的规范对应关系列表和默认的列表进行叠加，否则程序内置的规范对应关系将失效，只检查自定义的规范列表。设置了过滤条件的情况下，后续的处理过程会选择满足过滤条件的记录来进行，或者对于数据源的输入进行过滤后的记录可以直接持久化保存结果，如通过 qiniu 源获取文件列表过滤后进行保存，可设置 save-total=true/false 来选择是否将列举到的完整记录进行保存。")
	fmt.Println("filter 详细配置可见 filter 配置说明 https://github.com/NigelWu95/qiniu-suits-java/blob/master/docs/filter.md")
	fmt.Println()
	fmt.Println("过滤示例: ")
	fmt.Println("1、过滤 test 目录下 mime（content-type）为 video 类型的文件：")
	fmt.Println("f-prefix=test/")
	fmt.Println("f-mime=video/")
	fmt.Println()
	fmt.Println("2、过滤 2019 年 10 月 1 日 10 点之前的 .jpg 文件：")
	fmt.Println("f-suffix=.jpg")
	fmt.Println("f-date-scale=0,2019-10-01 10:00:00")
}
