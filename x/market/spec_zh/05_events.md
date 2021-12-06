<!--
市场 ：5
-->

# 事件

市场模块发出以下事件：

## 处理程序

### MsgSwap

|类型 |属性键 |属性值 |
|---------|---------------|--------------------|
| swap    | offer         | {offerCoin}        |
| swap    | trader        | {traderAddress}    |
| swap    | recipient     | {recipientAddress} |
| swap    | swap_coin     | {swapCoin}         |
| swap    | swap_fee      | {swapFee}          |
| message | module        | market             |
| message | action        | swap               |
| message | sender        | {senderAddress}    |

### MsgSwapSend

|类型 |属性键 |属性值 |
|---------|---------------|--------------------|
| swap    | offer         | {offerCoin}        |
| swap    | trader        | {traderAddress}    |
| swap    | recipient     | {recipientAddress} |
| swap    | swap_coin     | {swapCoin}         |
| swap    | swap_fee      | {swapFee}          |
| message | module        | market             |
| message | action        | swapsend           |
| message | sender        | {senderAddress}    |
