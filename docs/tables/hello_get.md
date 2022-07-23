# Table: hello_get

This table displays a hardcoded greeting.

## Examples

```
> select * from hello_get where id = '17'
+----+----------+-------------------+
| id | greeting | json              |
+----+----------+-------------------+
| 17 | Hello    | {"hello":"world"} |
+----+----------+-------------------+
```

