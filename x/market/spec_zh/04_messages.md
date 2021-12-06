<!--
市场：4
-->

# 消息

## MsgSwap

MsgSwap 交易表示交易者打算将他们的“OfferCoin”余额换成新面额的“AskDenom”，用于 Terra<>Terra 和 Terra<>Luna 交换。

```go
type MsgSwap struct {
	Trader    sdk.AccAddress
	OfferCoin sdk.Coin
	AskDenom  string
}
```

## MsgSwapSend
MsgSendSwap 首先将 OfferCoin 交换为 AskDenom，然后将结果币发送到 ToAddress。税是正常收取的，就好像发件人正在用交换的结果币发出 MsgSend 一样。


```go
type MsgSwapSend struct {
	FromAddress sdk.AccAddress
	ToAddress   sdk.AccAddress 
	OfferCoin   sdk.Coin
	AskDenom    string
}
```

## 功能

### ComputeSwap

```go
func (k Keeper) ComputeSwap(ctx sdk.Context, offerCoin sdk.Coin, askDenom string) (retDecCoin sdk.DecCoin, spread sdk.Dec, err error)
```

此函数从报价中检测掉期类型并询问面额并返回：

1. 给定的“offerCoin”应该返回的被询问的硬币数量。这是通过首先将“offerCoin”现货转换为 µSDR，然后使用 Oracle 报告的适当汇率从 µSDR 转换为所需的“askDenom”来实现的。

2. 给定掉期类型应被视为掉期费的点差百分比。 Terra<>Terra 掉期只包含托宾税点差费。 Terra<>Luna 点差是“MinSpread”和恒定产品定价的点差中的较大者。

如果 offerCoin 的面额与 `askDenom` 相同，这将引发 ErrRecursiveSwap。

### ApplySwapToPool

```go
func (k Keeper) ApplySwapToPool(ctx sdk.Context, offerCoin sdk.Coin, askCoin sdk.DecCoin) error
```

当 Terra 和 Luna 流动性池的余额发生变化时，此函数在交换期间被调用以更新区块链的度量，`TerraPoolDelta`。

Terra 货币共享相同的流动性池，因此在 Terra<>Terra 掉期期间，“TerraPoolDelta”保持不变。

对于 Terra<>Luna 交换，交换后池的相对大小会有所不同，并且 `delta` 将使用以下公式更新：

对于 Terra 到 Luna，`delta = delta + offerAmount`
对于 Luna 到 Terra，`delta = delta - askAmount` 