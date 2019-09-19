# go_study

demo

### interface
1. 森林结构，和java的树结构不一样
> 子对象可以付给父对象引用（interface的引用定义）
>> mike 和 tom 实现 Men
>> var i Men
>> i = mike
>> i = tom
2. 空的interface，类似于java中的Object，但比Object强大，可以代表一切
> 一个函数把interface{}作为参数，那么他可以接受任意类型的值作为参数，如果一个函数返回interface{},那么也就可以返回任意类型的值。


### 25个关键字
- var和const参考2.2Go语言基础里面的变量和常量申明
- package和import已经有过短暂的接触
- func 用于定义函数和方法
- return 用于从函数返回
- defer 用于类似析构函数
- go 用于并发
- select 用于选择不同类型的通讯
- interface 用于定义接口，参考2.6小节
- struct 用于定义抽象数据类型，参考2.5小节
- break、case、continue、for、fallthrough、else、if、switch、goto、default这些参考2.3流程介绍里面
- chan用于channel通讯
- type用于声明自定义类型
- map用于声明map类型数据
- range用于读取slice、map、channel数据