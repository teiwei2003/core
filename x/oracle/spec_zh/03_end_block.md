<!--
预言机：3
-->

# 结束区块

## 计算汇率投票

在每个块的末尾，“Oracle”模块检查它是否是“VotePeriod”的最后一个块。如果是，它运行[投票程序](./01_concepts.md#Voting_Procedure)：

1. 所有当前活跃的Luna汇率从商店中清除

2. 收到的选票按面额组织成选票。弃权投票，以及不活跃或被监禁的验证者的投票被忽略

3. 不符合以下要求的面额将被删除：

    - 必须出现在“白名单”中允许的面额中
    - 面额选票必须至少具有“VoteThreshold”总投票权

4. 对于每个剩余的带有通过选票的“denom”：

    - 统计选票并使用“tally()”找到加权中值汇率和获胜者
    - 遍历投票的获胜者并将他们的权重添加到他们的总和中
    - 使用 `k.SetLunaExchangeRate()` 在区块链上为 Luna<>`denom` 设置 Luna 汇率
   - 发出一个 `exchange_rate_update` 事件

5. 计算 [missed](./01_concepts.md#Slashing) Oracle 投票的验证器并增加相应的 Miss 计数器

6. 如果在`SlashWindow`的末尾，惩罚错过超过惩罚阈值的验证器（提交的有效票数少于`MinValidPerWindow`)

7. 使用 `k.RewardBallotWinners()` 向投票获胜者分发奖励

8. 清除所有prevotes（下一个`VotePeriod`除外)和来自商店的投票 