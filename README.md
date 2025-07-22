### Заменяем теги

```shell
sed -i 's|<number value="\([^"]*\)"></number>|<number value="\1"/>|g' contacts.xml
```
