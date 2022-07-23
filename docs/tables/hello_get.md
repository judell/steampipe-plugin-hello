# Table: hello_get

This tables defines a Get hydrate.

```
func tableHelloGet(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "hello_get",
		Description: "hello using Get/GetConfig",
		Get: &plugin.GetConfig{
			Hydrate: getGreeting,
			KeyColumns: plugin.SingleColumn("id"),
		},
		Columns: helloCols(),
	}
}
```

## Examples

```
> select * from hello_get

Error: rpc error: code = Internal desc = 'Get' call for table 'hello_get' is missing 1 required qual: column:'id' operator: = (SQLSTATE HV000)
```

> select * from hello_get where id = '17'
+----+----------+------+------+
| id | greeting | json | _ctx |
+----+----------+------+------+
+----+----------+------+------+
```

