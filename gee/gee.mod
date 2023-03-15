module example
go 1.19

require gee v0.0.0
replace gee => ./gee

<!--在go.mod中使用replace 将gee 指向 ./gee-->