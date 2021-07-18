# decorator

Go 语言的修饰器编程模式，也就是函数式编程(Functional Option)的模式。

## 基本概念

典型的例子是kubernetes client-go包中[NewSharedInformerFactoryWithOptions](https://github.com/kubernetes/client-go/blob/master/informers/factory.go#L109)这个函数。**options参数就是一些装饰器函数，在函数中遍历各个装饰器函数依次执行。在go中，这种做法更准确的叫做：Functional Options**

这种模式可以很轻松地把一些函数装配到另外一些函数上，让你的代码更加简单，也可以让一些“小功能型”的代码复用性更高，让代码中的函数可以像乐高玩具那样自由地拼装。但是go没有python和java那样的语法糖，用起来不是那么便捷。


