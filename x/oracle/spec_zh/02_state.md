<!--
预言机：2
-->

＃ 状态

## 汇率预先投票

`ExchangeRatePrevote` 包含验证者投票者对当前 `VotePeriod` 给定 denom 的预投票。

- ExchangeRatePrevote：`0x01<denom_Bytes><valAddress_Bytes> ->amino(ExchangeRatePrevote)`

```go
type ValAddress []byte
type VoteHash []byte

type ExchangeRatePrevote struct {
	Hash        VoteHash       // Vote hex hash to protect centralize data source problem
	Denom       string         // Ticker name of target fiat currency
	Voter       sdk.ValAddress // Voter val address
	SubmitBlock int64
}
```

## 汇率投票

`ExchangeRateVote` 包含验证者投票者对当前 `VotePeriod` 给定 denom 的投票。

- ExchangeRateVote：`0x02<denom_Bytes><valAddress_Bytes> ->amino(ExchangeRateVote)`

```go
type ExchangeRateVote struct {
	ExchangeRate sdk.Dec        // ExchangeRate of Luna in target fiat currency
	Denom        string         // Ticker name of target fiat currency
	Voter        sdk.ValAddress // voter val address of validator
}
```

## 汇率

一个 `sdk.Dec` 存储当前 Luna 对给定面值的汇率，[Market](../../market/spec/README.md) 模块使用它来为掉期定价。

您可以使用“k.GetActiveDenoms()”获取与“Luna”(票数超过“VoteThreshold”的面额)交易的活跃面额列表。

- ExchangeRate: `0x03<denom_Bytes> -> amino(sdk.Dec)`

## FeederDelegation

`operator` 的委托价格馈送器的 `sdk.AccAddress`(`terra-` 帐户)地址。

- FeederDelegation：`0x04<valAddress_Bytes> ->amino(sdk.AccAddress)`

## MissCounter

一个 `int64` 表示验证器 `operator` 在当前 `SlashWindow` 期间错过的 `VotePeriods` 的数量。

- MissCounter：`0x05<valAddress_Bytes> ->amino(int64)`

## AggregateExchangeRatePrevote

`AggregateExchangeRatePrevote` 包含验证者选民对当前 `VotePeriod` 的所有 denom 的聚合预投票。

- AggregateExchangeRatePrevote：`0x06<valAddress_Bytes> ->amino(AggregateExchangeRatePrevote)`

```go
// AggregateVoteHash 是隐藏投票汇率的哈希值
// 格式为 SHA256 中的十六进制字符串("{salt}:{exchange rate}{denom},...,{exchange rate}{denom}:{voter}")
type AggregateVoteHash []byte

type AggregateExchangeRatePrevote struct {
	Hash        AggregateVoteHash // Vote hex hash to protect centralize data source problem
	Voter       sdk.ValAddress    // Voter val address
	SubmitBlock int64
}
```

## AggregateExchangeRateVote

`AggregateExchangeRateVote` 包含验证者选民对当前 `VotePeriod` 的所有 denom 的总投票。

- AggregateExchangeRateVote:`0x07<valAddress_Bytes> ->amino(AggregateExchangeRateVote)`

```go
type ExchangeRateTuple struct {
	Denom        string  `json:"denom"`
	ExchangeRate sdk.Dec `json:"exchange_rate"`
}

type ExchangeRateTuples []ExchangeRateTuple

type AggregateExchangeRateVote struct {
	ExchangeRateTuples ExchangeRateTuples // ExchangeRates of Luna in target fiat currencies
	Voter              sdk.ValAddress     // voter val address of validator
}
```

## 托宾税

`sdk.Dec` 为投票通过的 denom 存储传播税，[Market](../../market/spec/README.md) 模块使用它来对 Terra<>Terra 进行现货转换。

- TobinTax：`0x08<denom_Bytes> ->amino(sdk.Dec)` 