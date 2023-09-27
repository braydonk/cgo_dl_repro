.PHONY: lib
liblib:
	gcc -c -Wall -Werror -fpic lib/lib.c
	gcc -shared -o liblib.so lib.o
