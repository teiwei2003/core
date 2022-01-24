<!--
国库：2
-->

# 状态

## 税率

当前时期的税率政策杠杆的价值。

- 税率：`0x01 -> amino(sdk.Dec)`

## 奖励权重

当前时期的奖励权重政策杠杆的值。

- RewardWeight：`0x02 -> amino(sdk.Dec)`

## 税收上限

财政部保留了一个“KVStore”，它将面额“denom”映射到一个“sdk.Int”，代表该面额交易的税收可以产生的最大收入。每个时期都会使用当前汇率下的“TaxPolicy.Cap”的等效值进行更新。

例如，如果交易价值为 100 SDT，税率和税收上限分别为 5% 和 1 SDT，则交易产生的收入将是 1 SDT 而不是 5 SDT，因为它超过了税收上限。

- TaxCap：`0x03<denom_Bytes> ->amino(sdk.Int)`

## 税收收入

当前时期的税收奖励 $T$。

- 税收收入：`0x04 -> amino(sdk.Coins)`

## 纪元初始发行量

当前纪元开始时 Luna 的总供应量。这个值在 `k.SettleSeigniorage()` 中用于计算在 epoch 结束时要分配的铸币税。

记录初始发行量会自动使用`Supply`模块来确定Luna的总发行量。为方便起见，Peeking 会将 epoch 最初发行的 µLuna 返回为 `sdk.Int` 而不是 `sdk.Coins`。

- EpochInitialIssuance: `0x05 -> amino(sdk.Coins)`

## 指标
财政部跟踪当前和以前时期的以下指标：

### 税收奖励
“epoch”的税收奖励。

- TaxReward：`0x06<epoch_Bytes> -> amino(sdk.Dec)`

### 铸币税奖励
“epoch”的铸币税奖励 $S$。

- SeigniorageReward：`0x07<epoch_Bytes> -> amino(sdk.Dec)`

### 总质押 Luna
`epoch` 的总质押 Luna $\lambda$。

- TotalStakedLuna: `0x08<epoch_Bytes> -> amino(sdk.Int)`

## 累积高度

将指标保持在硬分叉上的累积高度。

- CumulativeHeight：`0x09 -> amino(int64)` 