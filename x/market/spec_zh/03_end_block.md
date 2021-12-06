<!--
市场：3
-->

# 结束区块

## 补充池
在每个 `EndBlock` 处，`TerraPoolDelta` 的值根据参数的 `PoolRecoveryPeriod` 减少。

这允许网络在价格剧烈波动期间大幅提高点差费用，并在价格长期变化的一段时间后自动将点差恢复到正常水平。

```go
func (k Keeper) ReplenishPools(ctx sdk.Context) {
	delta := k.GetTerraPoolDelta(ctx)
	regressionAmt := delta.QuoInt64(k.PoolRecoveryPeriod(ctx))

	// Replenish terra pool towards base pool
	// regressionAmt cannot make delta zero
	delta = delta.Sub(regressionAmt)

	k.SetTerraPoolDelta(ctx, delta)
}
```
