<!--
国库：3
-->

# 结束块

如果区块链在 epoch 的最后一个区块，则运行以下过程：

1. 使用 `k.UpdateIndicators()` 更新所有指标

2. 如果当前区块在[probation](./01_concepts.md#Probation)下，跳到步骤6。

3. 结算 epoch 期间产生的铸币税，并将资金用于下一个 epoch 期间的投票奖励和社区池。

4. 计算下一个时期的“税率”、“奖励权重”和“税收上限”。

5. 发出`policy_update` 事件，记录新的政策杠杆值。

6. 最后，用`k.RecordEpochInitialIssuance()`记录Luna发行。这将用于计算下一个时期的铸币税。

# 功能

## `k.UpdateIndicators()`

```go
func (k Keeper) UpdateIndicators(ctx sdk.Context)
```

该函数在一个纪元结束时运行，并记录税收奖励 $T$、铸币税奖励 $S$ 和总抵押 Luna $\Sigma$ 的当前值作为纪元 $t$ 的历史指标，然后再移动到下一个纪元纪元 $t+1$。

$T_t$ 是 TaxProceeds 中的当前值
,$S_t = \Sigma * w$ 带有纪元铸币税 $\Sigma$ 和奖励权重 $w$。
$\lambda _t$ 只是 `staking.TotalBondedTokens()` 的结果。

## `k.UpdateIndicators()`

```go
func (k Keeper) UpdateTaxPolicy(ctx sdk.Context) (newTaxRate sdk.Dec)
```

该函数在一个时期结束时被调用，以计算税率货币杠杆的下一个值。

将 $\tau _t$ 视为当前税率，将 $n$ 视为 `MiningIncrement` 参数。

1. 计算过去一年`WindowLong`每单位Luna的税收奖励的滚动平均值$\tau_y$。

2. 计算上个月 `WindowShort` 上每单位 Luna 的税收奖励的滚动平均值 $\tau_m$`。

3. 如果$\tau _m = 0$，则上个月没有税收。因此，应根据“pc.Clamp()”的规则将税率设置为税收政策允许的最大值。

4. 否则，新的税率为 $r_{t+1} = n r_t \tau _y / \tau _m$，受`pc.Clamp()`的规则约束。

因此，当与长期税收平均水平相比，较短时间窗口内的税收表现不佳时，财政部会提高税率。当短期税收收入超过长期指数时，它会降低税率。

## `k.UpdateRewardPolicy()`

```go
func (k Keeper) UpdateRewardPolicy(ctx sdk.Context) (newRewardWeight sdk.Dec)
```

该函数在一个时期结束时被调用，以计算奖励权重货币杠杆的下一个值。

将 $w_t$ 视为当前奖励权重，将 $b$ 视为 SeigniorageBurdenTarget 参数。

1.计算上个月`WindowShort`的铸币税奖励总和$S_m$。

2.计算上个月`WindowShort`的挖矿总奖励$R_m$的总和。

3. 如果 $R_m = 0$ 或 $S_m = 0$，则上个月没有采矿和铸币税奖励。因此，应将奖励权重设置为奖励政策允许的最大值，并遵守“pc.Clamp()”的规则。

4. 否则，新的奖励权重为 $w_{t+1} = b w_t S_m / R_m$，受`pc.Clamp()`的规则约束。

### `k.UpdateTaxCap()`

```go
func (k Keeper) UpdateTaxCap(ctx sdk.Context) sdk.Coins
```

这个函数在一个时期结束时被调用，以计算下一个时期每个面额的税收上限。

对于流通中的每个面额，该面额的新税收上限设置为“TaxPolicy”参数中定义的全球税收上限，按当前汇率计算。

### `k.SettleSeigniorage()`

```go
func (k Keeper) SettleSeignioage(ctx sdk.Context)
```

这个函数在一个纪元结束时被调用来计算铸币税并将资金转发到 [`Oracle`](../../oracle/spec/README.md) 模块以获得投票奖励，以及 [`Distribution` ]（https://github.com/cosmos/cosmos-sdk/tree/master/x/distribution/spec/README.md）用于社区池。

1. 当前纪元的铸币税 $\Sigma$ 是通过取纪元开始时的 Luna 供应量 ([Epoch Initial Issuance](./02_state.md#EpochInitialIssuance)) 与当前纪元的 Luna 供应量之间的差值来计算的打电话的时间。

   请注意，当当前 Luna 供应量低于纪元开始时，$\Sigma > 0$，因为 Luna 已从 Luna 交换到 Terra 中被烧毁。见[这里](../../market/spec/01_concepts.md#Seigniorage)。

2. 奖励权重 $w$ 是指定用于投票奖励的铸币税的百分比。铸造了 $S$ 的新 Luna，并且 [`Oracle`](../../oracle/spec/README.md) 模块接收铸币税的 $S = \Sigma * w$。

3. 剩余的币 $\Sigma - S$ 发送到 [`Distribution`](https://github.com/cosmos/cosmos-sdk/tree/master/x/distribution/spec/README.md ) 模块，它被分配到社区池中。

## 政策约束
来自治理提案和自动校准的政策更新分别受“TaxPolicy”和“RewardPolicy”参数的约束。 “PolicyConstraints”类型指定了每个变量的下限、上限和最大周期性变化。

```go
// PolicyConstraints 定义关于更新关键财政部变量的约束
type PolicyConstraints struct {
    RateMin       sdk.Dec  `json:"rate_min"`
    RateMax       sdk.Dec  `json:"rate_max"`
    Cap           sdk.Coin `json:"cap"`
    ChangeRateMax sdk.Dec  `json:"change_max"`
}
```

约束策略杠杆更新的逻辑由`pc.Clamp()` 执行，如下所示。

```go
// 钳位在策略约束内约束策略变量更新
func (pc PolicyConstraints) Clamp(prevRate sdk.Dec, newRate sdk.Dec) (clampedRate sdk.Dec) {
	if newRate.LT(pc.RateMin) {
		newRate = pc.RateMin
	} else if newRate.GT(pc.RateMax) {
		newRate = pc.RateMax
	}

	delta := newRate.Sub(prevRate)
	if newRate.GT(prevRate) {
		if delta.GT(pc.ChangeRateMax) {
			newRate = prevRate.Add(pc.ChangeRateMax)
		}
	} else {
		if delta.Abs().GT(pc.ChangeRateMax) {
			newRate = prevRate.Sub(pc.ChangeRateMax)
		}
	}
	return newRate
}
``` 