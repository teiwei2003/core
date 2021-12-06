<!--
订单：5
-->

# 事件

oracle 模块发出以下事件：

## EndBlocker

|类型 |属性键 |属性值 |
|----------------------|---------------|-----------------|
| exchange_rate_update | denom         | {denom}         |
| exchange_rate_update | exchange_rate | {exchangeRate}  |  

## 处理程序

### MsgExchangeRatePrevote

|类型 |属性键 |属性值 |
|---------|---------------|---------------------|
| prevote | denom         | {denom}             |
| prevote | voter         | {validatorAddress}  |
| prevote | feeder        | {feederAddress}     |
| message | module        | oracle              |
| message | action        | exchangerateprevote |
| message | sender        | {senderAddress}     |

### MsgExchangeRateVote

|类型 |属性键 |属性值 |
|---------|---------------|--------------------|
| vote    | denom         | {denom}            |
| vote    | voter         | {validatorAddress} |
| vote    | exchange_rate | {exchangeRate}     |
| vote    | feeder        | {feederAddress}    |
| message | module        | oracle             |
| message | action        | exchangeratevote   |
| message | sender        | {senderAddress}    |


### MsgDelegateFeedConsent

|类型 |属性键 |属性值 |
|---------------|---------------|--------------------|
| feed_delegate | operator      | {validatorAddress} |
| feed_delegate | feeder        | {feederAddress}    |
| message       | module        | oracle             |
| message       | action        | delegatefeeder     |
| message       | sender        | {senderAddress}    |

### MsgAggregateExchangeRatePrevote

|类型 |属性键 |属性值 |
|-------------------|---------------|------------------------------|
| aggregate_prevote | voter         | {validatorAddress}           |
| aggregate_prevote | feeder        | {feederAddress}              |
| message           | module        | oracle                       |
| message           | action        | aggregateexchangerateprevote |
| message           | sender        | {senderAddress}              |

### MsgAggregateExchangeRateVote

|类型 |属性键 |属性值 |
|----------------|----------------|---------------------------|
| aggregate_vote | voter          | {validatorAddress}        |
| aggregate_vote | exchange_rates | {exchangeRates}           |
| aggregate_vote | feeder         | {feederAddress}           |
| message        | module         | oracle                    |
| message        | action         | aggregateexchangeratevote |
| message        | sender         | {senderAddress}           |
