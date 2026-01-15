# check-login

用 Bloom Filter 模拟“用户是否存在”的快速判断，并统计插入/命中/误判的耗时与误判率。

## 结构
- `bloom_filter.go`：加载用户 ID，构建 Bloom Filter，执行命中与未命中查询统计。
- `gen_ids.go`：生成模拟用户 ID 数据。
- `data/user_ids.txt`：默认输入数据（可自行生成替换）。

## 依赖
- Go（版本参考 `go.mod`）

## 使用
生成用户 ID 数据（默认 100 万条）：
```bash
go run gen_ids.go -count 1000000 -out data/user_ids.txt
```

运行 Bloom Filter 统计：
```bash
go run bloom_filter.go -in data/user_ids.txt -false-pos 0.001 -queries 100000
```

输出包含插入耗时、命中检查耗时、未命中检查耗时，以及误判比例。
