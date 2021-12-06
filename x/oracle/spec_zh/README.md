## 抽象的

Oracle 模块为 Terra 区块链提供了 Luna 与各种 Terra 挂钩汇率的最新且准确的价格馈送，以便 [Market](../../market/spec/README.md) 可以提供Terra<>Terra 货币对以及 Terra<>Luna 之间的公平交易。

由于价格信息是区块链外部的，Terra 网络依赖验证器定期对当前 Luna 汇率进行投票，协议每个“VotePeriod”对结果进行一次统计，并更新链上汇率作为加权中位数选票。

> 由于 Oracle 服务由验证器提供支持，您可能会发现 [Staking](https://github.com/cosmos/cosmos-sdk/tree/master/x/staking/spec/README.md ) 模块，它涵盖了 staking 和验证器的逻辑。

## 内容

1. **[概念](01_concepts.md)**
    - [投票程序](01_concepts.md#Voting-Procedure)
    - [奖励带](01_concepts.md#Reward-Band)
    - [削减](01_concepts.md#Slashing)
    - [弃权](01_concepts.md#Abstaining-from-Voting)
2. **[状态](02_state.md)**
    - [汇率预先投票](02_state.md#ExchangeRatePrevote)
    - [汇率投票](02_state.md#ExchangeRateVote)
    - [汇率](02_state.md#ExchangeRate)
    - [FeederDelegation](02_state.md#FeederDelegation)
    - [MissCounter](02_state.md#MissCounter)
    - [AggregateExchangeRatePrevote](02_state.md#AggregateExchangeRatePrevote)
    - [AggregateExchangeRateVote](02_state.md#AggregateExchangeRateVote)
    - [托宾税](02_state.md#TobinTax)
3. **[结束区块](03_end_block.md)**
    - [计算汇率投票](03_end_block.md#Tally-Exchange-Rate-Votes)
4. **[消息](04_messages.md)**
    - [MsgExchangeRatePrevote](04_messages.md#MsgExchangeRatePrevote)
    - [MsgExchangeRatePrevote](04_messages.md#MsgExchangeRatePrevote)
    - [MsgDelegateFeedConsent](04_messages.md#MsgDelegateFeedConsent)
    - [MsgAggregateExchangeRatePrevote](04_messages.md#MsgAggregateExchangeRatePrevote)
    - [MsgAggregateExchangeRateVote](04_messages.md#MsgAggregateExchangeRateVote)
5. **[事件](05_events.md)**
    - [EndBlocker](05_events.md#EndBlocker)
    - [处理程序](05_events.md#Handlers)
6. **[参数](06_params.md)** 