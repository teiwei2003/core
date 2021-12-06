## 摘要

Market 模块包含 Terra 货币之间（例如 UST<>KRT）以及 Terra 和 Luna（例如 SDT<>Luna）之间的原子交换逻辑。

以不同 Terra 面额之间以及 Terra 和 Luna 之间的公平汇率保证可用、流动性市场的能力对于用户采用和价格稳定至关重要。

TerraSDR 与 SDR 挂钩的价格稳定性是通过 Terra<>Luna 对协议的算法做市商的套利活动实现的，该活动扩大和收缩 Terra 供应以维持挂钩。

## 内容

1. **[概念](01_concepts.md)**
    - [交换费用](01_concepts.md#Swap-Fees)
    - [做市算法](01_concepts.md#Market-Making-Algorithm)
    - [虚拟流动性池](01_concepts.md#Virtual-Liquidity-Pools)
    - [交换程序](01_concepts.md#Swap-Procedure)
    - [铸币税](01_concepts.md#铸币税)
2. **[状态](02_state.md)**
    - [TerraPoolDelta](02_state.md#TerraPoolDelta)
3. **[结束区块](03_end_block.md)**
    - [补充池](03_end_block.md#Replenish-Pool)
4. **[消息](04_messages.md)**
    - [MsgSwap](04_messages.md#MsgSwap)
    - [MsgSwapSend](04_messages.md#MsgSwapSend)
    - [功能](04_messages.md#Functions)
5. **[事件](05_events.md)**
    - [处理程序](05_events.md#Handlers)
5. **[参数](06_params.md)** 