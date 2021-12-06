<!--
市场：2
-->

# 状态

## TerraPoolDelta

Market 模块提供基于恒定产品机制的交换功能。 Terra 矿池必须保持其 delta 以跟踪货币对交换价差的需求。 Luna 池可以通过以下等式从 Terra 池 delta 中检索：

```go
TerraPool := BasePool + delta
LunaPool := (BasePool * BasePool) / TerraPool
```

> 请注意，all 池中保存的是 `usdr` 金额的十进制单位，因此 delta 也是 `usdr` 单位。

- TerraPoolDelta：`0x01 -> amino(TerraPoolDelta)`

```go
type TerraPoolDelta sdk.Dec // TerraPool 和 BasePool 之间的差距
``` 