## KV Persistence Architecture

Underlying persistence is a KV store that is easy to replay state back onto

Keys are typically the type and then a hierarchical path to the value

### models
Key: type.<namespace>.<type>
Value: The representation of the data type

### Object
Key: object.<namespace>.<type>.<object>
Value: The binary value of the object

### Code
Key: code.<namespace>.<code>
Value: The source code in bytes value

### Accounts
An account can be returned by the address (signer) or by the friendly name

Key: signer.<namespace>.<address>
Key: account.<namespace>.<account>

Value: The account