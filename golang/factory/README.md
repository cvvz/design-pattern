# factory

## 基本概念

目的是**把具体对象的创建与业务代码解绑**。

业务代码不需要关心怎么出生产出一个具体的对象，也就是**不在业务代码里创建具体的对象**。

更通俗地说就是不在业务逻辑里new一个对象，而是统一通过一个`NewInstance`入口函数来生成不同的对象实例。

## 设计要点

1. **简单工厂**
    在go中，我们一般通过`NewInstance`来创建对象，只要返回值是接口，能根据不同的入参生成不同的对象实例，就是简单工厂。
