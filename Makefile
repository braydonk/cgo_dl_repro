.PHONY: lib
liblib:
	gcc -c -rdynamic -Wall -Werror -fpic lib/lib.c
	gcc -shared -o liblib.so lib.o

.PHONY: mainc
mainc:
	gcc -ggdb -Wl,--unresolved-symbols=ignore-in-object-files -o c/main c/main.c

.PHONY: mainasm
mainasm:
	gcc -S -Wl,--unresolved-symbols=ignore-in-object-files -o c/main.s c/main.c

.PHONY: goasm
goasm:
	go build -o cgo_dl_repro .
	go tool objdump cgo_dl_repro > gomain.s
