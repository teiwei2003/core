<!--
国库：4
-->

## 提案

财政部模块定义了特殊提案，允许对“KVStore”中的 [Tax Rate](./02_state.md#TaxRate) 和 [Reward Weight](./02_state.md#RewardWeight) 值进行投票并相应地更改，受`pc.Clamp()` 强加的[策略约束](./03_end_block.md#PolicyConstraints) 约束。

### 税率更新提案

```go
type TaxRateUpdateProposal struct {
	Title       string  // Title of the Proposal
	Description string  // Description of the Proposal
	TaxRate     sdk.Dec // target TaxRate
}
```

::: 详细信息 JSON 示例

```json
{
  "type": "treasury/TaxRateUpdateProposal",
  "value": {
    "title": "proposal title",
    "description": "proposal description",
    "tax_rate": "0.001000000000000000"
  }
}
```

### 奖励权重更新提案

```go
type RewardWeightUpdateProposal struct {
	Title        string  // Title of the Proposal
	Description  string  // Description of the Proposal
	RewardWeight sdk.Dec // target RewardWeight
}
```

::: 详细信息 JSON 示例

```json
{
  "type": "treasury/RewardWeightUpdateProposal",
  "value": {
    "title": "proposal title",
    "description": "proposal description",
    "reward_weight": "0.001000000000000000"
  }
}
```