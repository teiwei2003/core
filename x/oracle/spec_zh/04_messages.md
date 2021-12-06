<!--
预言机：4
-->

# 消息

## MsgExchangeRatePrevote(已弃用)

`Hash` 是由格式为`{salt}:{exchange_rate}:{denom}:{voter}` 的字符串的 SHA256 哈希(十六进制字符串)的前 20 个字节生成的十六进制字符串，实际的元数据`MsgExchangeRateVote` 将在下一个 `VotePeriod` 中跟进。你可以使用`GetVoteHash()` 函数来帮助编码这个哈希。请注意，由于在随后的“MsgExchangeRateVote”中，必须显示盐，因此必须为每次预投票提交重新生成使用的盐。

“Denom”是投票所针对的货币的面额。例如，如果选民希望为美元提交预投票，那么正确的“Denom”是“uusd”。

哈希中使用的汇率必须是 Luna 的公开市场汇率，相对于与“Denom”匹配的面额。例如，如果`Denom` 是`uusd` 并且Luna 的现行汇率是1 美元，那么必须使用“1”作为汇率，因为`1 uluna = 1 uusd`。

`Feeder`(`terra-` 地址)用于如果验证者希望将 oracle 投票签名委托给一个单独的密钥(代替运营商“提供”价格)以降低暴露其验证者签名密钥的风险。

`Validator` 是原始验证器的验证器地址(`terravaloper-` 地址)。


```go
// 已弃用：在 columbus-4 之后将弃用正常的 prevote 和投票
// MsgExchangeRatePrevote - 用于对 ExchangeRateVote 进行投票的结构。
// prevote的目的是用hash隐藏投票汇率
// 在 SHA256("{salt}:{exchange_rate}:{denom}:{voter}") 中格式化为十六进制字符串
type MsgExchangeRatePrevote struct {
	Hash      VoteHash
	Denom     string
	Feeder    sdk.AccAddress
	Validator sdk.ValAddress
}
```

## MsgExchangeRateVote(已弃用)

`MsgExchangeRateVote` 包含实际的汇率投票。 `Salt` 参数必须与用于创建预投票的盐相匹配，否则无法奖励投票者。

```go
// 已弃用：在 columbus-4 之后将弃用正常的 prevote 和投票
// MsgExchangeRateVote - 用于对以各种 Terra 资产计价的 Luna 汇率进行投票的结构。
// 例如，如果验证者认为 Luna 的美元有效汇率是 10.39，那就是
// 汇率字段是什么，如果 KRW 为 1213.34，则相同。
type MsgExchangeRateVote struct {
	ExchangeRate sdk.Dec        // the effective rate of Luna in {Denom}
	Salt         string
	Denom        string
	Feeder       sdk.AccAddress
	Validator    sdk.ValAddress
}
```

## MsgDelegateFeedConsent

验证者还可以选择将投票权委托给另一个密钥，以防止块签名密钥保持在线。为此，他们必须提交“MsgDelegateFeedConsent”，将他们的预言机投票权委托给代表验证者签署“MsgExchangeRatePrevote”和“MsgExchangeRateVote”的“Delegate”。

> 委托验证者可能会要求您存入一些资金(在 Terra 或 Luna 中)，他们可以用这些资金来支付费用，并通过单独的 MsgSend 发送。该协议是在链下制定的，不受 Terra 协议的强制执行。

`Operator` 字段包含验证器的操作员地址(前缀为 `terravaloper-`)。 `Delegate` 字段是将代表 `Operator` 提交与汇率相关的投票和预投票的受托人账户的账户地址(前缀为 `terra-`)。

```go
// MsgDelegateFeedConsent - 将预言机投票权委托给另一个地址的结构。
type MsgDelegateFeedConsent struct {
	Operator sdk.ValAddress 
	Delegate sdk.AccAddress 
}
```

## MsgAggregateExchangeRatePrevote

`Hash` 是由格式为 `{salt}:{exchange rate}{denom},...,{exchange rate}{denom} 的字符串的 SHA256 哈希(十六进制字符串)的前 20 个字节生成的十六进制字符串}:{voter}`，下一个 `VotePeriod` 中要遵循的实际 `MsgAggregateExchangeRateVote` 的元数据。你可以使用`GetAggregateVoteHash()` 函数来帮助编码这个散列。请注意，由于在随后的“MsgAggregateExchangeRateVote”中，必须显示盐，因此必须为每次预投票提交重新生成使用的盐。

```go
// MsgAggregateExchangeRatePrevote - 用于对 ExchangeRateVote 进行聚合预投票的结构。
//聚合prevote的目的是用hash隐藏投票汇率
// 格式为 SHA256 中的十六进制字符串("{salt}:{exchange rate}{denom},...,{exchange rate}{denom}:{voter}")
type MsgAggregateExchangeRatePrevote struct {
	Hash      AggregateVoteHash 
	Feeder    sdk.AccAddress    
	Validator sdk.ValAddress    
}
```

## MsgAggregateExchangeRateVote

`MsgAggregateExchangeRateVote` 包含实际汇率投票。 `Salt` 参数必须与用于创建预投票的盐相匹配，否则无法奖励投票者。

```go
// MsgAggregateExchangeRateVote - 用于对以各种 Terra 资产计价的 Luna 汇率进行投票的结构。
type MsgAggregateExchangeRateVote struct {
	Salt          string
	ExchangeRates string
	Feeder        sdk.AccAddress 
	Validator     sdk.ValAddress 
}
```
