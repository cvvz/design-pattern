# builder

## 解决问题

1. 避免初始化类时，需要传入太多的参数。比如必填项太多时，虽然可以通过set函数来减少参数传入，但是就没法校验必填项是不是都填了。
2. 没法进行配置项之间的依赖关系或约束条件的校验，除非让所有的配置项都以参数的形式传进去。

## 基本概念

把校验逻辑放置到 Builder 类中，先创建建造者，并且通过 set() 方法设置建造者的变量值，然后在使用 build() 方法真正创建对象之前，做集中的校验，校验通过之后才会创建对象。

[kubectl](https://github.com/kubernetes/kubectl/blob/master/pkg/cmd/describe/describe.go#L155-L165)源码中使用到了Builder这个设计模式。

## 副作用

**使用建造者模式来构建对象，代码实际上是有点重复的**，需要多加一个Builder类，并且 Builder 类中的属性实际和被build出来的对象中的属性一样。在golang中，可以使用[Functional Options](../decorator/decorator.go)，以函数式编程的思想来解决。
