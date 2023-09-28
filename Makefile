.PHONY: lib
liblib:
	gcc -c -rdynamic -Wall -Werror -fpic lib/lib.c
	gcc -shared -o liblib.so lib.o

.PHONY: mainc
mainc:
	gcc -Wl,--unresolved-symbols=ignore-in-object-files -o main main.c
