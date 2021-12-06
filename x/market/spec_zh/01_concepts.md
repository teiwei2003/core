<!--
市场：1
-->

# 概念

## 交换费
由于 Terra 的价格馈送源自验证者预言机，因此链上报告的价​​格与实际实时价格之间必然存在延迟。

这种差异大约为 1 分钟(我们的预言机 VotePeriod 为 30 秒)，这对于几乎所有实际交易都可以忽略不计。然而，攻击者可以利用这种滞后并通过抢先攻击从网络中提取价值。

为了防止这种情况，市场模块强制执行以下交换费用

* 用于现货转换 Terra<>Terra 交换的托宾税(设置为 0.25%)

    举例说明，假设 oracle 报告 Luna<>SDT 汇率为 10，而 Luna<>KRT 汇率为 10,000。发送 1 SDT 将获得 0.1 Luna，即 1000 KRT。应用托宾税后，您最终会得到 997.5 KRT(1000 的 0.25% 是 2.5)，这比任何零售货币兑换和汇款都要优惠。

* Terra<>Luna 交换的最小点差(设置为 2%)

    使用上述相同的汇率，交换 1 SDT 将返回价值 980 KRT 的 Luna(1000 的 2% 为 20，作为交换费用)。另一方面，1 Luna 会给你 9.8 SDT(10 的 2% = 0.2)或 9800 KRT(10,000 的 2% = 200)。

## 做市算法
Terra 使用恒定乘积做市算法来确保 Terra<>Luna 交换的流动性。

对于 Constant Product，我们定义了一个值“CP”，该值设置为 Terra 池的大小乘以 Luna 的设定法定值，并确保我们的做市商通过调整点差在任何交换期间保持不变。

> 注意 - 我们对 Constant Product 的实现与 Uniswap 不同，因为我们使用 Luna 的法定价值而不是 Luna 池的大小。这种细微差别意味着 Luna 价格的变化不会影响产品，而是影响 Luna 池的大小。

```
CP = TerraPool * LunaPool * LunaPice / SDRPrice
```

例如，我们将从相等的 Terra 和 Luna 池开始，两者的总价值为 1000 SDR。 Terra 矿池的大小为 1000 SDT，假设 Luna<>SDR 的价格为 0.5，则 Luna 矿池的大小为 2000 Luna。用 100 SDT 交换 Luna 将返回价值约 90.91 SDR 的 Luna(≈ 181.82 Luna)。 100 SDT的offer加入Terra矿池，价值90.91 SDT的Luna从Luna矿池中取出。

```
CP = 1000000 SDR
(1000 SDT) * (Luna 的 1000 SDR) = 1000000 SDR
(1100 SDT) * (909.0909... Luna SDR) = 1000000 SDR
```

当然，这个具体的例子是为了说明而不是现实——生产中使用的流动性池要大得多，差价的幅度就会减小。

与 Columbus-2 相比，Constant-Product 的主要优势在于它提供了“无限”的流动性，从某种意义上说，可以提供任意规模的交换服务(尽管价格随着交易规模的增加而变得越来越不利)。

## 虚拟流动资金池

市场从两个大小相等的流动性池开始，一个代表 Terra(所有面额)，另一个代表 Luna，由参数“BasePool”初始化，该参数定义了 Terra 和 Luna 流动性池的初始大小。

在实践中，不是跟踪两个池的大小，而是将信息编码为一个数字“delta”，区块链将其存储为“TerraPoolDelta”，代表 Terra 池与其基本大小的偏差，单位为 µSDR。

Terra 和 Luna 流动性池的大小可以使用以下公式生成：

```
TerraPool = BasePool + delta
LunaPool * LunaPice / SDRPrice = (BasePool * BasePool) / TerraPool
LunaPool = (SDRPrice / LunaPrice) * (BasePool * BasePool) / TerraPool
```

在每个区块结束时，市场模块将尝试通过降低 Terra 和 Luna 池之间的大小来“补充”池。池的补充率由参数“PoolRecoveryPeriod”设置，周期越短意味着对交易的敏感性越低，这意味着之前的交易被更快地遗忘，市场能够提供更多的流动性。

这种机制确保流动性并充当一种低通滤波器，允许点差费用(这是 TerraPoolDelta 的函数)在需求发生变化时回落，因此需要吸收必要的供应变化.

## 交换程序

1. Market 模块接收 `MsgSwap` 消息并执行基本的验证检查

2. 使用`k.ComputeSwap()`计算`ask`和`spread`

3. 使用 `k.ApplySwapToPool()` 更新 `TerraPoolDelta`

4. 使用 `supply.SendCoinsFromAccountToModule()` 将 `OfferCoin` 从账户转移到模块

5. 使用 `supply.BurnCoins()` 销毁提供的硬币。

6. 让`fee = spread * ask`，这是点差费用。

7. 使用 `supply.MintCoins()` 铸造 `AskDenom` 的 `ask - fee` 硬币。当“费用”硬币被销毁时，这隐含地应用了点差费用。

8. 使用 `supply.SendCoinsFromModuleToAccount()` 将新铸造的硬币发送给交易者

9. 发出`swap`事件来宣传交换并记录点差费

如果交易者的“账户”余额不足以执行交换，则交换交易失败。

成功完成 Terra<>Luna 交换后，将记入用户账户的部分代币作为点差费用被扣留。

## 铸币税
Luna 交换到 Terra 时，被协议重新捕获的 Luna 被烧毁，称为铸币税——发行新 Terra 产生的价值。 在 epoch 结束时，将计算该 epoch 的总铸币税并将其重新引入经济中作为汇率预言机的投票奖励和由财政部模块引入社区池，[此处](../. ./treasury/spec/README.md)。