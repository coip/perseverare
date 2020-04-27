usage: replace the `${tokens}` below to vet

``` bash
$ go get https://github.com/coip/perseverare
$  username=${userID} \
    password=${password} \
    host=${servername}.mariadb.database.azure.com \
    default_schema=${default_schema} \
    perseverare
```

the leading space in above snippet is intentional*, will (~typically) allow performing this check **without** logging the related secrets to ~bash_history.

also, you do need the mariadb in azure of course. this is already relatively banal, but without that, pfft.

*:[whole reason i made this md file]