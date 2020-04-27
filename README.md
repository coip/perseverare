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

----- 

edit:

seems to also require at-least github.com/go-sql-driver/mysql@v1.4, perhaps even v1.5. (a toml w/ v1.3 in-part got me here)
> src: https://github.com/go-sql-driver/mysql/commit/9181e3a86a19bacd63e68d43ae8b7b36320d8092#diff-eeb2e1d2b086ffdc5e8a6e60178702aa
