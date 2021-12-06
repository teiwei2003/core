# 事件

## 处理程序

## MsgStoreCode

|类型 |属性键 |属性值 |
| ---------- | ------------- | --------------- |
| store_code | sender        | {senderAddress} |
| store_code | codeID        | {codeID}        |
| message    | module        | wasm            |
| message    | action        | store_code      |
| message    | sender        | {senderAddress} |

## MsgMigrateCode

|类型 |属性键 |属性值 |
| ------------ | ------------- | --------------- |
| migrate_code | sender        | {senderAddress} |
| migrate_code | codeID        | {codeID}        |
| message      | module        | wasm            |
| message      | action        | migrate_code    |
| message      | sender        | {senderAddress} |

## MsgInstantiateContract

|类型 |属性键 |属性值 |
| -------------------- | ---------------- | -------------------- |
| instantiate_contract | creator          | {creatorAddress}     |
| instantiate_contract | admin            | {adminAddress}       |
| instantiate_contract | code_id          | {codeID}             |
| instantiate_contract | contract_address | {contractAddress}    |
| message              | module           | wasm                 |
| message              | action           | instantiate_contract |
| message              | sender           | {senderAddress}      |

## MsgExecuteContract

|类型 |属性键 |属性值 |
| ---------------- | ---------------- | ----------------- |
| execute_contract | sender           | {senderAddress}   |
| execute_contract | contract_address | {contractAddress} |
| wasm-*           | ...              | ...               |
| wasm             | ...              | ...               |
| from_contract    | ...              | ...               |
| message          | module           | wasm              |
| message          | action           | execute_contract  |
| message          | sender           | {senderAddress}   |

## MsgMigrateContract

|类型 |属性键 |属性值 |
| ---------------- | ---------------- | ----------------- |
| migrate_contract | code_id          | {codeID}          |
| migrate_contract | contract_address | {contractAddress} |
| message          | module           | wasm              |
| message          | action           | migrate_contract  |
| message          | sender           | {senderAddress}   |

## MsgUpdateContractAdmin

|类型 |属性键 |属性值 |
| --------------------- | ---------------- | --------------------- |
| update_contract_admin | admin            | {adminAddress}        |
| update_contract_admin | contract_address | {contractAddress}     |
| message               | module           | wasm                  |
| message               | action           | update_contract_admin |
| message               | sender           | {senderAddress}       |

## MsgClearContractAdmin

|类型 |属性键 |属性值 |
| -------------------- | ---------------- | -------------------- |
| clear_contract_admin | contract_address | {contractAddress}    |
| message              | module           | wasm                 |
| message              | action           | clear_contract_admin |
| message              | sender           | {senderAddress}      |
